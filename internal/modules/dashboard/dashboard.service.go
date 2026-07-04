package dashboard

type DashboardService interface {
	GetPerformanceOverview(year int) (PerformanceOverviewResponse, error)
}
