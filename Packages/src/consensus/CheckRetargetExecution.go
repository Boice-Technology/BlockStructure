package consensus

func CheckRetargetExecution(height int64, startIntervalTimestamp int64, endIntervalTimestamp int64, currentTargetDifficulty string) (string, bool){
	var newTarget string = ""
	var verdict bool = false
	if height % 2016 == 0 {
		timeInterval := endIntervalTimestamp - startIntervalTimestamp
		newTarget = RetargetDifficulty(currentTargetDifficulty, timeInterval)
		verdict = true
	}
	return newTarget, verdict
}