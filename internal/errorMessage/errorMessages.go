package errorMessage

// Developer
const InvalidDeveloperNameErrorMessage = "developer name can not be empty"
const InvalidDeveloperDedicationErrorMessage = "developer dedication must be positive"
const InvalidDeveloperRateErrorMessage = "developer rate per hour must be positive"

// Team
const InvalidTeamResponsibleErrorMessage = "team must be composed of subteams or developers and these can not be duplicated"
const InvalidTeamNameErrorMessage = "team name can not be empty"

// ConcreteTask
const InvalidConcreteTaskNameErrorMessage = "concrete task name can not be empty"
const InvalidConcreteTaskEffortErrorMessage = "concrete task effort must be positive"
const InvalidConcreteTaskDependentsErrorMessage = "concrete task can not have direct repeated tasks"

// Project
const InvalidProjectNameErrorMessage = "project name can not be empty"
const InvalidProjectSubtasksErrorMessage = "project must have at least one subtask"
