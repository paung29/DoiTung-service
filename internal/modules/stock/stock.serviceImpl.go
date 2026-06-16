package stock

import (
	"errors"

	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/modules/customer"
	"github.com/doitung/DoiTung-service/internal/modules/warehouse"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/types/enums"
	"github.com/doitung/DoiTung-service/internal/utils"
	"gorm.io/gorm"
)

type service struct {
	db            *gorm.DB
	repo          StockRepository
	yearRepo      year.YearRepository
	warehouseRepo warehouse.WarehouseRepository
	customerRepo  customer.CustomerRepository
}

func NewStockService(db *gorm.DB, repo StockRepository, yearRepo year.YearRepository, warehouseRepo warehouse.WarehouseRepository, customerRepo customer.CustomerRepository) StockService {
	return &service{db: db, repo: repo, yearRepo: yearRepo, warehouseRepo: warehouseRepo, customerRepo: customerRepo}
}

func (s *service) CreateCarryOver(accountID uint, form CreateCarryOverStockRequest) (StockMovementResponse, error) {

	yearRecord, err := s.yearRepo.FindByYear(int(form.Year))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Year doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve year record")
	}
	yearId := yearRecord.YearID

	ProductYearRecord, err := s.yearRepo.FindByYear(int(*form.ProductionYear))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Production year doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve production year record")
	}
	productionYearId := ProductYearRecord.YearID

	warehouseRecord, err := s.warehouseRepo.FindByID(*form.WarehouseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Warehouse doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve warehouse record")
	}
	warehouseId := warehouseRecord.WarehouseID

	if form.TotalGrams == nil && form.TotalPods == nil {
		return StockMovementResponse{}, utils.BadRequestError("Total grams or total pods must be provided")
	}

	// transaction start here
	tx := s.db.Begin()
	if tx.Error != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to start database transaction")
	}

	// Create stock movement record if fail rollback transaction and return error
	stockMovement := &models.StockMovement{
		RecordedByID:     accountID,
		YearID:           yearId,
		ProductionYearID: &productionYearId,
		ToWarehouseID:    &warehouseId,
		Grade:            form.Grade,
		TotalGrams:       form.TotalGrams,
		TotalPods:        form.TotalPods,
		Details:          form.Details,
		RecordedDate:     form.RecordedDate,
		MovementType:     enums.MovementCarryOver,
	}
	err = s.repo.CreateStockMovement(tx, stockMovement)
	if err != nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.SystemError("Failed to create stock movement")
	}

	stockBalanceRecord, err := s.repo.GetStockBalanceForUpdate(tx, productionYearId, warehouseId, form.Grade)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If no existing stock balance record, create a new one with the carry over quantity
			stockBalanceRecord = &models.StockBalance{
				YearID:      productionYearId,
				WarehouseID: warehouseId,
				Grade:       form.Grade,
				TotalGrams:  0,
				TotalPods:   0,
			}
			err = s.repo.CreateNewStockBalance(tx, stockBalanceRecord)
			if err != nil {
				tx.Rollback()
				return StockMovementResponse{}, utils.SystemError("Failed to create new stock balance")
			}
		} else {
			tx.Rollback()
			return StockMovementResponse{}, utils.SystemError("Failed to retrieve stock balance")
		}
	}

	// Update stock balance by adding the carry over quantity to the existing balance
	if form.TotalGrams != nil {
		stockBalanceRecord.TotalGrams += *form.TotalGrams
	}
	if form.TotalPods != nil {
		stockBalanceRecord.TotalPods += *form.TotalPods
	}
	err = s.repo.UpdateStockBalance(tx, stockBalanceRecord)
	if err != nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.SystemError("Failed to update stock balance")
	}
	if err = tx.Commit().Error; err != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to commit database transaction")
	}

	return StockMovementResponse{Message: "Stock movement created successfully"}, nil
}

func (s *service) CreateIncomingStock(accountID uint, form CreateIncomingStockRequest) (StockMovementResponse, error) {

	yearRecord, err := s.yearRepo.FindByYear(int(form.Year))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Year doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve year record")
	}

	ProductYearRecord, err := s.yearRepo.FindByYear(int(*form.ProductionYear))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Production year doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve production year record")
	}

	warehouseRecord, err := s.warehouseRepo.FindByID(*form.WarehouseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Warehouse doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve warehouse record")
	}

	if form.TotalGrams == nil && form.TotalPods == nil {
		return StockMovementResponse{}, utils.BadRequestError("Total grams or total pods must be provided")
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to start database transaction")
	}
	stockMovement := &models.StockMovement{
		RecordedByID:     accountID,
		YearID:           yearRecord.YearID,
		ProductionYearID: &ProductYearRecord.YearID,
		ToWarehouseID:    &warehouseRecord.WarehouseID,
		Grade:            form.Grade,
		TotalGrams:       form.TotalGrams,
		TotalPods:        form.TotalPods,
		Details:          form.Details,
		RecordedDate:     form.RecordedDate,
		MovementType:     enums.MovementIncoming,
	}
	err = s.repo.CreateStockMovement(tx, stockMovement)
	if err != nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.SystemError("Failed to create stock movement")
	}

	stockBalanceRecord, err := s.repo.GetStockBalanceForUpdate(tx, ProductYearRecord.YearID, warehouseRecord.WarehouseID, form.Grade)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stockBalanceRecord = &models.StockBalance{
				YearID:      ProductYearRecord.YearID,
				WarehouseID: warehouseRecord.WarehouseID,
				Grade:       form.Grade,
				TotalGrams:  0,
				TotalPods:   0,
			}
			err = s.repo.CreateNewStockBalance(tx, stockBalanceRecord)
			if err != nil {
				tx.Rollback()
				return StockMovementResponse{}, utils.SystemError("Failed to create new stock balance")
			}
		} else {
			tx.Rollback()
			return StockMovementResponse{}, utils.SystemError("Failed to retrieve stock balance")
		}
	}

	if form.TotalGrams != nil {
		stockBalanceRecord.TotalGrams += *form.TotalGrams
	}
	if form.TotalPods != nil {
		stockBalanceRecord.TotalPods += *form.TotalPods
	}
	err = s.repo.UpdateStockBalance(tx, stockBalanceRecord)
	if err != nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.SystemError("Failed to update stock balance")
	}
	if err = tx.Commit().Error; err != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to commit database transaction")
	}

	return StockMovementResponse{Message: "Stock movement created successfully"}, nil
}

func (s *service) CreateIssuedStock(accountID uint, form CreateIssuedStockRequest) (StockMovementResponse, error) {

	yearRecord, err := s.yearRepo.FindByYear(int(form.Year))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Year doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve year record")
	}

	ProductYearRecord, err := s.yearRepo.FindByYear(int(*form.ProductionYear))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Production year doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve production year record")
	}
	productionYearId := ProductYearRecord.YearID

	warehouseRecord, err := s.warehouseRepo.FindByID(*form.WarehouseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Warehouse doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve warehouse record")
	}
	warehouseId := warehouseRecord.WarehouseID

	customerRecord, err := s.customerRepo.FindByCustomerID(*form.CustomerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Customer doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve customer record")
	}

	if int(form.TotalGrams) <= 0 && int(form.TotalPods) <= 0 {
		return StockMovementResponse{}, utils.BadRequestError("Total grams or total pods must be greater than 0")
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to start database transaction")
	}
	// Get total stock for the specified production year, warehouse, and grade
	stockBalance, err := s.repo.GetStockBalanceForUpdate(tx, productionYearId, warehouseId, form.Grade)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("No stock available for the specified production year, warehouse, and grade")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve stock balance")
	}

	if stockBalance == nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.BadRequestError("No stock available for the specified production year, warehouse, and grade")
	}

	if stockBalance.TotalGrams == 0 && stockBalance.TotalPods == 0 {
		tx.Rollback()
		return StockMovementResponse{}, utils.BadRequestError("No stock available for the specified production year, warehouse, and grade")
	}

	// If the total stock is less than the requested issued quantity, return an error
	if stockBalance.TotalGrams < int(form.TotalGrams) || stockBalance.TotalPods < int(form.TotalPods) {
		tx.Rollback()
		return StockMovementResponse{}, utils.BadRequestError("Insufficient stock available for the requested issue quantity")
	}

	stockMovement := &models.StockMovement{
		RecordedByID:       accountID,
		YearID:             yearRecord.YearID,
		ProductionYearID:   &productionYearId,
		FromWarehouseID:    &warehouseId,
		IssuedToCustomerID: &customerRecord.CustomerID,
		Grade:              form.Grade,
		PricePerGram:       &form.PricePerGram,
		TotalGrams:         &form.TotalGrams,
		TotalPods:          &form.TotalPods,
		Details:            form.Details,
		RecordedDate:       form.RecordedDate,
		MovementType:       enums.MovementIssued,
	}
	err = s.repo.CreateStockMovement(tx, stockMovement)
	if err != nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.SystemError("Failed to create stock movement")
	}

	// Update stock balance by subtracting the issued quantity from the existing balance
	stockBalance.TotalGrams -= int(form.TotalGrams)
	stockBalance.TotalPods -= int(form.TotalPods)
	err = s.repo.UpdateStockBalance(tx, stockBalance)
	if err != nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.SystemError("Failed to update stock balance")
	}

	err = tx.Commit().Error
	if err != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to commit database transaction")
	}

	return StockMovementResponse{Message: "Stock movement created successfully"}, nil
}

func (s *service) DeleteStockMovement(stockMovementID uint) (StockMovementResponse, error) {
	stockMovementRecord, err := s.repo.FindByID(stockMovementID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockMovementResponse{}, utils.BadRequestError("Stock movement record doesn't exist")
		}
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve stock movement record")
	}

	tx := s.db.Begin()

	if tx.Error != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to start database transaction")
	}

	isIssuedMovement := stockMovementRecord.MovementType == enums.MovementIssued

	if stockMovementRecord.ProductionYearID == nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.BadRequestError("Stock movement record has no production year")
	}
	if stockMovementRecord.Grade == "" {
		tx.Rollback()
		return StockMovementResponse{}, utils.BadRequestError("Stock movement record has no grade")
	}

	if stockMovementRecord.TotalGrams == nil || stockMovementRecord.TotalPods == nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.BadRequestError("Stock movement record has no quantity information")
	}
	productionYearId := *stockMovementRecord.ProductionYearID
	grade := stockMovementRecord.Grade
	warehouseId := uint(0)
	if isIssuedMovement {
		if stockMovementRecord.FromWarehouseID == nil {
			tx.Rollback()
			return StockMovementResponse{}, utils.BadRequestError("Stock movement record has no from warehouse information")
		}
		warehouseId = *stockMovementRecord.FromWarehouseID
	} else {
		if stockMovementRecord.ToWarehouseID == nil {
			tx.Rollback()
			return StockMovementResponse{}, utils.BadRequestError("Stock movement record has no to warehouse information")
		}
		warehouseId = *stockMovementRecord.ToWarehouseID
	}

	stockBalanceRecord, err := s.repo.GetStockBalanceForUpdate(tx, productionYearId, warehouseId, grade)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return StockMovementResponse{}, utils.BadRequestError("Stock balance record doesn't exist")
		}
		tx.Rollback()
		return StockMovementResponse{}, utils.SystemError("Failed to retrieve stock balance")
	}

	if isIssuedMovement {
		stockBalanceRecord.TotalGrams += *stockMovementRecord.TotalGrams
		stockBalanceRecord.TotalPods += *stockMovementRecord.TotalPods

	} else {
		stockBalanceRecord.TotalGrams -= *stockMovementRecord.TotalGrams
		stockBalanceRecord.TotalPods -= *stockMovementRecord.TotalPods
		if stockBalanceRecord.TotalGrams < 0 || stockBalanceRecord.TotalPods < 0 {
			tx.Rollback()
			return StockMovementResponse{}, utils.BadRequestError("Cannot delete this stock movement because some stock has already been issued")
		}
	}

	err = s.repo.UpdateStockBalance(tx, stockBalanceRecord)
	if err != nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.SystemError("Failed to update stock balance")
	}

	err = s.repo.DeleteStockMovement(tx, stockMovementRecord.StockMovementID)
	if err != nil {
		tx.Rollback()
		return StockMovementResponse{}, utils.SystemError("Failed to delete stock movement")
	}

	err = tx.Commit().Error
	if err != nil {
		return StockMovementResponse{}, utils.SystemError("Failed to commit database transaction")
	}

	return StockMovementResponse{Message: "Stock movement deleted successfully"}, nil
}

func (s *service) GetStockMovementListsByYear(year uint) (GetAllStockMovementsByYearResponse, error) {

	yearRecord, err := s.yearRepo.FindByYear(int(year))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return GetAllStockMovementsByYearResponse{}, utils.BadRequestError("Year doesn't exist")
		}
		return GetAllStockMovementsByYearResponse{}, utils.SystemError("Failed to retrieve year record")
	}

	stockMovements, err := s.repo.GetAllByYearId(yearRecord.YearID)
	if err != nil {
		return GetAllStockMovementsByYearResponse{}, utils.SystemError("Failed to retrieve stock movements")
	}

	var stockMovementDetailsList []StockMovementDetails
	for number, movement := range stockMovements {

		isIssuedMovement := movement.MovementType == enums.MovementIssued
		warehouseName := ""
		if isIssuedMovement {
			if movement.FromWarehouse != nil {
				warehouseName = movement.FromWarehouse.WarehouseName
			}
		} else {
			if movement.ToWarehouse != nil {
				warehouseName = movement.ToWarehouse.WarehouseName
			}
		}

		productionYear := 0
		if movement.ProductionYear != nil {
			productionYear = movement.ProductionYear.Year
		}

		totalGrams := 0
		if movement.TotalGrams != nil {
			totalGrams = *movement.TotalGrams
		}

		totalPods := 0
		if movement.TotalPods != nil {
			totalPods = *movement.TotalPods
		}

		details := StockMovementDetails{
			No:              uint(number + 1),
			StockMovementID: movement.StockMovementID,
			Date:            movement.RecordedDate.Format("2006-01-02"),
			Category:        movement.MovementType,
			Grade:           movement.Grade,
			ProductionYear:  productionYear,
			Warehouse:       warehouseName,
			TotalGrams:      totalGrams,
			TotalPods:       totalPods,
			Details:         movement.Details,
		}

		stockMovementDetailsList = append(stockMovementDetailsList, details)
	}
	return GetAllStockMovementsByYearResponse{StockMovements: stockMovementDetailsList}, nil
}

func (s *service) GetCustomerStockTableByYear(year uint) (CustomerStockTableByYearResponse, error) {

	yearRecord, err := s.yearRepo.FindByYear(int(year))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return CustomerStockTableByYearResponse{}, utils.BadRequestError("Year doesn't exist")
		}
		return CustomerStockTableByYearResponse{}, utils.SystemError("Failed to retrieve year record")
	}

	customerStockTable, err := s.repo.GetCustomerStockByYearId(yearRecord.YearID)
	if err != nil {
		return CustomerStockTableByYearResponse{}, utils.SystemError("Failed to retrieve customer stock table")
	}

	customerStockTableResponse := make([]CustomerStockTableItem, 0, len(customerStockTable))
	for number, stock := range customerStockTable {

		stocks := CustomerStockTableItem{
			CustomerID:   stock.CustomerId,
			No:           number + 1,
			CustomerName: stock.CustomerName,
			GradeA:       stock.GradeA,
			GradeB:       stock.GradeB,
			GradeC:       stock.GradeC,
			GradeFailed:  stock.GradeFailed,
			TotalWeight:  stock.TotalWeight,
			Note:         stock.Note,
		}
		customerStockTableResponse = append(customerStockTableResponse, stocks)

	}
	return CustomerStockTableByYearResponse{CustomerStockTable: customerStockTableResponse}, nil
}

func (s *service) GetStockOverviewBalanceByYear(year uint) (StockOverviewResponse, error) {

	yearRecord, err := s.yearRepo.FindByYear(int(year))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return StockOverviewResponse{}, utils.BadRequestError("Year doesn't exist")
		}
		return StockOverviewResponse{}, utils.SystemError("Failed to retrieve year record")
	}
	yearId := yearRecord.YearID

	stockBalance, err := s.repo.GetStockOverviewBalanceByYearId(yearId)
	if err != nil {
		return StockOverviewResponse{}, utils.SystemError("Failed to retrieve stock overview balance")
	}

	incomingStockBalance, err := s.repo.GetIncomingStockTotal(yearId)
	if err != nil {
		return StockOverviewResponse{}, utils.SystemError("Failed to retrieve incoming stock total")
	}

	issuedStockBalance, err := s.repo.GetIssuedStockTotal(yearId)
	if err != nil {
		return StockOverviewResponse{}, utils.SystemError("Failed to retrieve issued stock total")
	}

	monthlySummaryRecords, err := s.repo.GetMonthlySummary(yearId)
	if err != nil {
		return StockOverviewResponse{}, utils.SystemError("Failed to retrieve monthly summary")
	}

	totalPodInStock := 0
	totalGramInStock := 0

	for _, balance := range stockBalance {
		totalPodInStock += balance.TotalPods
		totalGramInStock += balance.TotalGrams
	}

	gradeMap := make(map[enums.Grade]GradeSummary)
	for _, grade := range stockBalance {
		gradeMap[grade.Grade] = grade
	}

	grades := []enums.Grade{
		enums.GradeAPlus,
		enums.GradeA,
		enums.GradeB,
		enums.GradeC,
		enums.GradeDPlus,
		enums.GradeD,
	}

	gradeSummary := make([]GradeSummaryItem, 0, len(grades))
	for _, grade := range grades {
		row := gradeMap[grade]

		percentage := 0.0
		if totalPodInStock > 0 {
			percentage = float64(row.TotalPods) * 100 / float64(totalPodInStock)
		}

		gradeSummary = append(gradeSummary, GradeSummaryItem{
			Grade:      grade,
			TotalPod:   row.TotalPods,
			TotalGram:  row.TotalGrams,
			TotalKg:    float64(row.TotalGrams) / 1000,
			Percentage: percentage,
		})
	}

	monthNames := []string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}

	monthlyMap := make(map[int]MonthlySummary)
	for _, monthly := range monthlySummaryRecords {
		monthlyMap[monthly.Month] = monthly
	}

	monthlySummary := make([]MonthlySummaryItem, 0, 12)
	runningTotalWeight := 0

	for month := 1; month <= 12; month++ {
		row := monthlyMap[month]
		runningTotalWeight += row.TotalWeight

		monthlySummary = append(monthlySummary, MonthlySummaryItem{
			Month:          month,
			MonthName:      monthNames[month-1],
			StockInWeight:  row.StockInWeight,
			StockOutWeight: row.StockOutWeight,
			TotalWeight:    runningTotalWeight,
		})
	}
	return StockOverviewResponse{
		TotalPodInStock:  totalPodInStock,
		TotalGramInStock: totalGramInStock,
		TotalKgInStock:   float64(totalGramInStock) / 1000,

		IncomingStockPod:  incomingStockBalance.TotalPods,
		IncomingStockGram: incomingStockBalance.TotalGrams,
		IncomingStockKg:   float64(incomingStockBalance.TotalGrams) / 1000,

		IssuedStockPod:  issuedStockBalance.TotalPods,
		IssuedStockGram: issuedStockBalance.TotalGrams,
		IssuedStockKg:   float64(issuedStockBalance.TotalGrams) / 1000,

		GradeSummary:   gradeSummary,
		MonthlySummary: monthlySummary,
	}, nil
}
