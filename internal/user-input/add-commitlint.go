package userinput

func AddCommitlint() bool {
	return ConfirmationHandler("Do you wanna add commitlint?:(y,n) ")
}
