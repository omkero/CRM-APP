package router

import (
	"crm_system/internal/constants"
	"crm_system/internal/controllers"
	"crm_system/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	Controller *controllers.Controllers
}

func NewRouter(Controller *controllers.Controllers) *Router {
	return &Router{
		Controller: Controller,
	}
}

func (R *Router) Routes(app *fiber.App) {
	var BASE_NAME string = constants.API_BASE_NAME

	{
		customer_endpoint := app.Group(BASE_NAME + "/customer")
		customer_endpoint.Get("/get_customer", middlewares.VerifyEmployeeAuthToken, R.Controller.GetCustomerInfo)
		customer_endpoint.Post("/create_customer", middlewares.VerifyEmployeeAuthToken, R.Controller.CreateNewCustomer)
		customer_endpoint.Delete("/delete_customer", middlewares.VerifyEmployeeAuthToken, R.Controller.DeleteCustomer)
		customer_endpoint.Get("/get_all_customers", middlewares.VerifyEmployeeAuthToken, R.Controller.GetAllCustomers)
	}
	{
		employee_endpoint := app.Group(BASE_NAME + "/employee")
		employee_endpoint.Post("/sigunp", middlewares.VerifyEmployeeAuthToken, R.Controller.SignUpEmployee) // protected handler
		employee_endpoint.Post("/signin", R.Controller.SignInEmployee)
		employee_endpoint.Get("/get_employee", middlewares.VerifyEmployeeAuthToken, R.Controller.GetEmployeeInformation)
		employee_endpoint.Get("/get_all_employees/:pageNum", middlewares.VerifyEmployeeAuthToken, R.Controller.GetAllEmployeesInformation)
		employee_endpoint.Post("/change_employee_position/", middlewares.VerifyEmployeeAuthToken, R.Controller.ChangeEmployeePosition)
		employee_endpoint.Post("/suspend_employee/", middlewares.VerifyEmployeeAuthToken, R.Controller.SuspendEmployee)
		employee_endpoint.Post("/cancel_suspension/", middlewares.VerifyEmployeeAuthToken, R.Controller.CancelSuspension)
		employee_endpoint.Post("/ban_employee/", middlewares.VerifyEmployeeAuthToken, R.Controller.BanEmployee)
		employee_endpoint.Post("/un_ban_employee/", middlewares.VerifyEmployeeAuthToken, R.Controller.UnBanEmployee)
	}
	{
		tasks_endpoint := app.Group(BASE_NAME + "/tasks")
		tasks_endpoint.Post("/create_task_to", middlewares.VerifyEmployeeAuthToken, R.Controller.CreateTaskToEmployee)
		tasks_endpoint.Get("/get_employee_tasks", middlewares.VerifyEmployeeAuthToken, R.Controller.GetEmployeeTasks)
		tasks_endpoint.Post("/delete_task", middlewares.VerifyEmployeeAuthToken, R.Controller.DeleteTask)
		tasks_endpoint.Post("/apply_task", middlewares.VerifyEmployeeAuthToken, R.Controller.ApplyTask)
		tasks_endpoint.Get("/get_all_tasks", middlewares.VerifyEmployeeAuthToken, R.Controller.GetAllTasks)
	}
	{
		role_endpoint := app.Group(BASE_NAME + "/role")
		role_endpoint.Post("/create_role", middlewares.VerifyEmployeeAuthToken, R.Controller.CreateRole)
		role_endpoint.Post("/update_role", middlewares.VerifyEmployeeAuthToken, R.Controller.UpdateRole)
		role_endpoint.Get("/get_system_roles", middlewares.VerifyEmployeeAuthToken, R.Controller.GetSystemRoles)
		role_endpoint.Post("/apply_roles", middlewares.VerifyEmployeeAuthToken, R.Controller.ApplyRole)
		role_endpoint.Get("/get_system_permissions", middlewares.VerifyEmployeeAuthToken, R.Controller.GetSystemPermissions)
	}
	{
		product_endpoint := app.Group(BASE_NAME + "/product")
		product_endpoint.Post("/create_product", middlewares.VerifyEmployeeAuthToken, R.Controller.CreateProduct)
		product_endpoint.Post("/delete_product", middlewares.VerifyEmployeeAuthToken, R.Controller.DeleteProduct)
		product_endpoint.Post("/update_product", middlewares.VerifyEmployeeAuthToken, R.Controller.UpdateProduct)
		product_endpoint.Get("/get_product", middlewares.VerifyEmployeeAuthToken, R.Controller.GetProduct)
		product_endpoint.Get("/get_all_products", middlewares.VerifyEmployeeAuthToken, R.Controller.GetAllProducts)
	}
}
