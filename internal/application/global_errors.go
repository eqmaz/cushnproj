package application

var errorMap = map[string]string{
	"eNcF01": "No valid cfg file found",
	"eDupNs": "Database user or password not set. Ensure they are set via environment variables.",
	"ePuNf1": "Invalid product UUID; product not found",
	"eFuNf1": "Invalid fund UUID; fund not found",
	"eAcEx1": "Customer already has a retail ISA account",
	"eAcEx2": "Account does not belong to the user",
	"eIafId": "Insufficient allowance for ISA deposit",
}
