package screen

type ScreenAction string

const (
	ActionRead ScreenAction = "read"

	ActionPress   ScreenAction = "press"
	ActionRelease ScreenAction = "release"

	ActionToggle ScreenAction = "toggle"
)

func (s ScreenAction) IsValid() bool {
	for _, a := range ScreenActions() {
		if a == s {
			return true
		}
	}

	return false
}

func ScreenActions() []ScreenAction {
	return []ScreenAction{
		ActionRead,
		ActionPress,
		ActionRelease,
		ActionToggle,
	}
}
