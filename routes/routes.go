package routes

import (
	"banking_system/controllers"
	"banking_system/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	bankService := services.NewBankService(db)
	branchService := services.NewBranchService(db)
	customerService := services.NewCustomerService(db)
	accountService := services.NewAccountService(db)
	loanService := services.NewLoanService(db)
	repaymentService := services.NewRepaymentService(db)
	transactionService := services.NewTransactionService(db)

	bankController := controllers.NewBankController(bankService)
	branchController := controllers.NewBranchController(branchService)
	customerController := controllers.NewCustomerController(customerService)
	accountController := controllers.NewAccountController(accountService)
	loanController := controllers.NewLoanController(loanService)
	repaymentController := controllers.NewRepaymentController(repaymentService)
	transactionController := controllers.NewTransactionController(transactionService)

	banks := router.Group("/banks")
	{
		banks.POST("", bankController.CreateBank)
		banks.GET("", bankController.GetAllBanks)
		banks.GET("/:id", bankController.GetBankByID)
		banks.PUT("/:id", bankController.UpdateBank)
		banks.DELETE("/:id", bankController.DeleteBank)
	}

	branches := router.Group("/branches")
	{
		branches.POST("", branchController.CreateBranch)
		branches.GET("", branchController.GetAllBranches)
		branches.GET("/:id", branchController.GetBranchByID)
		branches.PUT("/:id", branchController.UpdateBranch)
		branches.DELETE("/:id", branchController.DeleteBranch)
	}

	customers := router.Group("/customers")
	{
		customers.POST("", customerController.CreateCustomer)
		customers.GET("", customerController.GetAllCustomers)
		customers.GET("/:id", customerController.GetCustomerByID)
		customers.PUT("/:id", customerController.UpdateCustomer)
		customers.DELETE("/:id", customerController.DeleteCustomer)

		customers.GET("/:id/accounts", customerController.GetCustomerAccounts)
		customers.GET("/:id/loans", customerController.GetCustomerLoans)
	}

	accounts := router.Group("/accounts")
	{
		accounts.POST("", accountController.CreateAccount)
		accounts.GET("", accountController.GetAllAccounts)
		accounts.GET("/:id", accountController.GetAccountByID)
		accounts.PUT("/:id", accountController.UpdateAccount)
		accounts.DELETE("/:id", accountController.DeleteAccount)

		accounts.POST("/:id/customers/:customerId", accountController.AddCustomerToAccount)
		accounts.DELETE("/:id/customers/:customerId", accountController.RemoveCustomerFromAccount)

		accounts.GET("/:id/transactions", accountController.GetAccountTransactions)

		accounts.POST("/:id/deposit", accountController.Deposit)
		accounts.POST("/:id/withdraw", accountController.Withdraw)
	}

	loans := router.Group("/loans")
	{
		loans.POST("", loanController.CreateLoan)
		loans.GET("", loanController.GetAllLoans)
		loans.GET("/:id", loanController.GetLoanByID)
		loans.PUT("/:id", loanController.UpdateLoan)
		loans.DELETE("/:id", loanController.DeleteLoan)

		loans.GET("/:id/details", loanController.GetLoanDetails)
		loans.POST("/:id/repay", loanController.RepayLoan)
	}

	repayments := router.Group("/repayments")
	{
		repayments.POST("", repaymentController.CreateRepayment)
		repayments.GET("", repaymentController.GetAllRepayments)
		repayments.GET("/:id", repaymentController.GetRepaymentByID)
		repayments.PUT("/:id", repaymentController.UpdateRepayment)
		repayments.DELETE("/:id", repaymentController.DeleteRepayment)
	}

	transactions := router.Group("/transactions")
	{
		transactions.POST("", transactionController.CreateTransaction)
		transactions.GET("", transactionController.GetAllTransactions)
		transactions.GET("/:id", transactionController.GetTransactionByID)
		transactions.PUT("/:id", transactionController.UpdateTransaction)
		transactions.DELETE("/:id", transactionController.DeleteTransaction)
	}
	return router
}

