package bgp

const KILL_SESSION = 0

type Signal struct {
	SignalType uint8
}

func KillSession() Signal {
	return Signal{SignalType: KILL_SESSION}
}
