package types

// SEE_MASK
type SEE_MASK ULONG

const (
	SEE_MASK_DEFAULT            SEE_MASK = 0x00000000
	SEE_MASK_CLASSNAME          SEE_MASK = 0x00000001
	SEE_MASK_CLASSKEY           SEE_MASK = 0x00000003
	SEE_MASK_IDLIST             SEE_MASK = 0x00000004
	SEE_MASK_INVOKEIDLIST       SEE_MASK = 0x0000000c
	SEE_MASK_ICON               SEE_MASK = 0x00000010
	SEE_MASK_HOTKEY             SEE_MASK = 0x00000020
	SEE_MASK_NOCLOSEPROCESS     SEE_MASK = 0x00000040
	SEE_MASK_CONNECTNETDRV      SEE_MASK = 0x00000080
	SEE_MASK_NOASYNC            SEE_MASK = 0x00000100
	SEE_MASK_DOENVSUBST         SEE_MASK = 0x00000200
	SEE_MASK_FLAG_NO_UI         SEE_MASK = 0x00000400
	SEE_MASK_UNICODE            SEE_MASK = 0x00004000
	SEE_MASK_NO_CONSOLE         SEE_MASK = 0x00008000
	SEE_MASK_ASYNCOK            SEE_MASK = 0x00100000
	SEE_MASK_HMONITOR           SEE_MASK = 0x00200000
	SEE_MASK_NOZONECHECKS       SEE_MASK = 0x00800000
	SEE_MASK_NOQUERYCLASSSTORE  SEE_MASK = 0x01000000
	SEE_MASK_WAITFORINPUTIDLE   SEE_MASK = 0x02000000
	SEE_MASK_FLAG_LOG_USAGE     SEE_MASK = 0x04000000
	SEE_MASK_FLAG_HINST_IS_SITE SEE_MASK = 0x08000000
)

// SEE_HINSTANCE
type SEE_HINSTANCE UINT_PTR

const (
	SE_ERR_FNF             SEE_HINSTANCE = 2
	SE_ERR_PNF             SEE_HINSTANCE = 3
	SE_ERR_ACCESSDENIED    SEE_HINSTANCE = 5
	SE_ERR_OOM             SEE_HINSTANCE = 8
	SE_ERR_DLLNOTFOUND     SEE_HINSTANCE = 32
	SE_ERR_SHARE           SEE_HINSTANCE = 26
	SE_ERR_ASSOCINCOMPLETE SEE_HINSTANCE = 27
	SE_ERR_DDETIMEOUT      SEE_HINSTANCE = 28
	SE_ERR_DDEFAIL         SEE_HINSTANCE = 29
	SE_ERR_DDEBUSY         SEE_HINSTANCE = 30
	SE_ERR_NOASSOC         SEE_HINSTANCE = 31
)

type SHELLEXECUTEINFO struct {
	Size       DWORD
	Mask       SEE_MASK
	Hwnd       HWND
	Verb       LPCWSTR
	File       LPCWSTR
	Parameters LPCWSTR
	Directory  LPCWSTR
	Show       ShowWindowCmd
	InstApp    SEE_HINSTANCE
	IDList     LPVOID
	Class      LPCWSTR
	KeyClass   HKEY
	HotKey     DWORD
	Icon       HANDLE
	Process    HANDLE
}

type LPSHELLEXECUTEINFO *SHELLEXECUTEINFO
