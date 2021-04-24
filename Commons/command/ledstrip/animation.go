package ledstrip

type LedstripAnimation string

const (
	AnimationRead  LedstripAnimation = "read"
	AnimationWrite LedstripAnimation = "write"

	AnimationSprinkle   LedstripAnimation = "sprinkle"
	AnimationFlushRight LedstripAnimation = "flush_right"
	AnimationFlushLeft  LedstripAnimation = "flush_left"
)

func (s LedstripAnimation) IsValid() bool {
	for _, a := range ScreenActions() {
		if a == s {
			return true
		}
	}

	return false
}

func ScreenActions() []LedstripAnimation {
	return []LedstripAnimation{
		AnimationRead,
		AnimationWrite,
		AnimationSprinkle,
		AnimationFlushLeft,
		AnimationFlushRight,
	}
}
