package userinput

func ShouldPublishToNPM() bool {
	return ConfirmationHandler("Will you publish it on npm?:(y,n) ")
}
