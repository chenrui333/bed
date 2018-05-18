package editor

import (
	"github.com/itchyny/bed/event"
	"github.com/itchyny/bed/key"
	"github.com/itchyny/bed/mode"
)

func defaultKeyManagers() map[mode.Mode]*key.Manager {
	kms := make(map[mode.Mode]*key.Manager)
	km := key.NewManager(true)
	km.Register(event.Quit, "Z", "Q")
	km.Register(event.Quit, "c-w", "q")
	km.Register(event.Quit, "c-w", "c-q")
	km.Register(event.Quit, "c-w", "c")
	km.Register(event.Suspend, "c-z")

	km.Register(event.CursorUp, "up")
	km.Register(event.CursorDown, "down")
	km.Register(event.CursorLeft, "left")
	km.Register(event.CursorRight, "right")
	km.Register(event.PageUp, "pgup")
	km.Register(event.PageDown, "pgdn")
	km.Register(event.PageTop, "home")
	km.Register(event.PageEnd, "end")
	km.Register(event.CursorUp, "k")
	km.Register(event.CursorDown, "j")
	km.Register(event.CursorLeft, "h")
	km.Register(event.CursorRight, "l")
	km.Register(event.CursorPrev, "b")
	km.Register(event.CursorPrev, "backspace")
	km.Register(event.CursorPrev, "backspace2")
	km.Register(event.CursorNext, "w")
	km.Register(event.CursorNext, " ")
	km.Register(event.CursorHead, "0")
	km.Register(event.CursorHead, "^")
	km.Register(event.CursorEnd, "$")
	km.Register(event.ScrollUp, "c-y")
	km.Register(event.ScrollDown, "c-e")
	km.Register(event.ScrollTop, "z", "t")
	km.Register(event.ScrollTopHead, "z", "enter")
	km.Register(event.ScrollMiddle, "z", "z")
	km.Register(event.ScrollMiddleHead, "z", ".")
	km.Register(event.ScrollBottom, "z", "b")
	km.Register(event.ScrollBottomHead, "z", "-")
	km.Register(event.PageUp, "c-b")
	km.Register(event.PageDown, "c-f")
	km.Register(event.PageUpHalf, "c-u")
	km.Register(event.PageDownHalf, "c-d")
	km.Register(event.PageTop, "g", "g")
	km.Register(event.PageEnd, "G")
	km.Register(event.WindowTop, "H")
	km.Register(event.WindowMiddle, "M")
	km.Register(event.WindowBottom, "L")
	km.Register(event.JumpTo, "\x1d")
	km.Register(event.JumpBack, "c-t")
	km.Register(event.DeleteByte, "x")
	km.Register(event.DeleteByte, "delete")
	km.Register(event.DeletePrevByte, "X")
	km.Register(event.Increment, "c-a")
	km.Register(event.Increment, "+")
	km.Register(event.Decrement, "c-x")
	km.Register(event.Decrement, "-")
	km.Register(event.ShowBinary, "g", "b")

	km.Register(event.Paste, "p")
	km.Register(event.PastePrev, "P")

	km.Register(event.StartInsert, "i")
	km.Register(event.StartInsertHead, "I")
	km.Register(event.StartAppend, "a")
	km.Register(event.StartAppendEnd, "A")
	km.Register(event.StartReplaceByte, "r")
	km.Register(event.StartReplace, "R")

	km.Register(event.Undo, "u")
	km.Register(event.Redo, "c-r")

	km.Register(event.StartVisual, "v")

	km.Register(event.SwitchFocus, "tab")
	km.Register(event.SwitchFocus, "backtab")
	km.Register(event.StartCmdlineCommand, ":")
	km.Register(event.StartCmdlineSearchForward, "/")
	km.Register(event.StartCmdlineSearchBackward, "?")
	km.Register(event.NextSearch, "n")
	km.Register(event.PreviousSearch, "N")

	km.Register(event.New, "c-w", "n")
	km.Register(event.New, "c-w", "c-n")
	km.Register(event.FocusWindowDown, "c-w", "down")
	km.Register(event.FocusWindowDown, "c-w", "c-j")
	km.Register(event.FocusWindowDown, "c-w", "j")
	km.Register(event.FocusWindowUp, "c-w", "up")
	km.Register(event.FocusWindowUp, "c-w", "c-k")
	km.Register(event.FocusWindowUp, "c-w", "k")
	km.Register(event.FocusWindowLeft, "c-w", "left")
	km.Register(event.FocusWindowLeft, "c-w", "c-h")
	km.Register(event.FocusWindowLeft, "c-w", "backspace")
	km.Register(event.FocusWindowLeft, "c-w", "h")
	km.Register(event.FocusWindowRight, "c-w", "right")
	km.Register(event.FocusWindowRight, "c-w", "c-l")
	km.Register(event.FocusWindowRight, "c-w", "l")
	km.Register(event.FocusWindowTopLeft, "c-w", "t")
	km.Register(event.FocusWindowTopLeft, "c-w", "c-t")
	km.Register(event.FocusWindowBottomRight, "c-w", "b")
	km.Register(event.FocusWindowBottomRight, "c-w", "c-b")
	km.Register(event.FocusWindowPrevious, "c-w", "p")
	km.Register(event.FocusWindowPrevious, "c-w", "c-p")
	km.Register(event.MoveWindowTop, "c-w", "K")
	km.Register(event.MoveWindowBottom, "c-w", "J")
	km.Register(event.MoveWindowLeft, "c-w", "H")
	km.Register(event.MoveWindowRight, "c-w", "L")
	kms[mode.Normal] = km

	km = key.NewManager(false)
	km.Register(event.ExitInsert, "escape")
	km.Register(event.ExitInsert, "c-c")
	km.Register(event.CursorUp, "up")
	km.Register(event.CursorDown, "down")
	km.Register(event.CursorLeft, "left")
	km.Register(event.CursorRight, "right")
	km.Register(event.CursorUp, "c-p")
	km.Register(event.CursorDown, "c-n")
	km.Register(event.CursorPrev, "c-b")
	km.Register(event.CursorNext, "c-f")
	km.Register(event.PageUp, "pgup")
	km.Register(event.PageDown, "pgdn")
	km.Register(event.PageTop, "home")
	km.Register(event.PageEnd, "end")
	km.Register(event.Backspace, "backspace")
	km.Register(event.Backspace, "backspace2")
	km.Register(event.Delete, "delete")
	km.Register(event.SwitchFocus, "tab")
	km.Register(event.SwitchFocus, "backtab")
	kms[mode.Insert] = km
	kms[mode.Replace] = km

	km = key.NewManager(true)
	km.Register(event.ExitVisual, "escape")
	km.Register(event.ExitVisual, "c-c")
	km.Register(event.ExitVisual, "v")
	km.Register(event.SwitchVisualEnd, "o")
	km.Register(event.SwitchVisualEnd, "O")
	km.Register(event.StartCmdlineCommand, ":")

	km.Register(event.CursorUp, "up")
	km.Register(event.CursorDown, "down")
	km.Register(event.CursorLeft, "left")
	km.Register(event.CursorRight, "right")
	km.Register(event.PageUp, "pgup")
	km.Register(event.PageDown, "pgdn")
	km.Register(event.PageTop, "home")
	km.Register(event.PageEnd, "end")
	km.Register(event.CursorUp, "k")
	km.Register(event.CursorDown, "j")
	km.Register(event.CursorLeft, "h")
	km.Register(event.CursorRight, "l")
	km.Register(event.CursorPrev, "b")
	km.Register(event.CursorPrev, "backspace")
	km.Register(event.CursorPrev, "backspace2")
	km.Register(event.CursorNext, "w")
	km.Register(event.CursorNext, " ")
	km.Register(event.CursorHead, "0")
	km.Register(event.CursorHead, "^")
	km.Register(event.CursorEnd, "$")
	km.Register(event.ScrollUp, "c-y")
	km.Register(event.ScrollDown, "c-e")
	km.Register(event.ScrollTop, "z", "t")
	km.Register(event.PageUp, "c-b")
	km.Register(event.PageDown, "c-f")
	km.Register(event.PageUpHalf, "c-u")
	km.Register(event.PageDownHalf, "c-d")
	km.Register(event.PageTop, "g", "g")
	km.Register(event.PageEnd, "G")
	km.Register(event.SwitchFocus, "tab")
	km.Register(event.SwitchFocus, "backtab")
	km.Register(event.Copy, "y")
	km.Register(event.Cut, "x")
	km.Register(event.Cut, "d")
	km.Register(event.Cut, "delete")
	kms[mode.Visual] = km

	km = key.NewManager(false)
	km.Register(event.CursorLeft, "left")
	km.Register(event.CursorLeft, "c-b")
	km.Register(event.CursorRight, "right")
	km.Register(event.CursorRight, "c-f")
	km.Register(event.CursorHead, "home")
	km.Register(event.CursorHead, "c-a")
	km.Register(event.CursorEnd, "end")
	km.Register(event.CursorEnd, "c-e")
	km.Register(event.BackspaceCmdline, "c-h")
	km.Register(event.BackspaceCmdline, "backspace")
	km.Register(event.BackspaceCmdline, "backspace2")
	km.Register(event.DeleteCmdline, "delete")
	km.Register(event.DeleteWordCmdline, "c-w")
	km.Register(event.ClearToHeadCmdline, "c-u")
	km.Register(event.ClearCmdline, "c-k")
	km.Register(event.ExitCmdline, "escape")
	km.Register(event.ExitCmdline, "c-c")
	km.Register(event.CompleteForwardCmdline, "tab")
	km.Register(event.CompleteBackCmdline, "backtab")
	km.Register(event.ExecuteCmdline, "enter")
	km.Register(event.ExecuteCmdline, "c-j")
	km.Register(event.ExecuteCmdline, "c-m")
	kms[mode.Cmdline] = km
	kms[mode.Search] = km
	return kms
}
