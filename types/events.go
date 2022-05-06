package types

// Variables
type (
	PEVENT_CALLBACK              LPVOID
	PEVENT_TRACE_BUFFER_CALLBACK LPVOID
	PEVENT_RECORD_CALLBACK       LPVOID
)

// WNODE_HEADER_u1_s
type WNODE_HEADER_u1_s struct {
	Version ULONG
	Linkage ULONG
}

// WNODE_HEADER_u1
type WNODE_HEADER_u1 struct {
	HistoricalContext ULONG64
}

// WNODE_HEADER_u2
type WNODE_HEADER_u2 struct {
	CountLost ULONG
}

// WNODE_HEADER_Flags
type WNODE_HEADER_Flags ULONG

const (
	WNODE_FLAG_ALL_DATA              WNODE_HEADER_Flags = 0x00000001
	WNODE_FLAG_SINGLE_INSTANCE       WNODE_HEADER_Flags = 0x00000002
	WNODE_FLAG_SINGLE_ITEM           WNODE_HEADER_Flags = 0x00000004
	WNODE_FLAG_EVENT_ITEM            WNODE_HEADER_Flags = 0x00000008
	WNODE_FLAG_FIXED_INSTANCE_SIZE   WNODE_HEADER_Flags = 0x00000010
	WNODE_FLAG_TOO_SMALL             WNODE_HEADER_Flags = 0x00000020
	WNODE_FLAG_INSTANCES_SAME        WNODE_HEADER_Flags = 0x00000040
	WNODE_FLAG_STATIC_INSTANCE_NAMES WNODE_HEADER_Flags = 0x00000080
	WNODE_FLAG_INTERNAL              WNODE_HEADER_Flags = 0x00000100
	WNODE_FLAG_USE_TIMESTAMP         WNODE_HEADER_Flags = 0x00000200
	WNODE_FLAG_PERSIST_EVENT         WNODE_HEADER_Flags = 0x00000400
	WNODE_FLAG_EVENT_REFERENCE       WNODE_HEADER_Flags = 0x00002000
	WNODE_FLAG_ANSI_INSTANCENAMES    WNODE_HEADER_Flags = 0x00004000
	WNODE_FLAG_METHOD_ITEM           WNODE_HEADER_Flags = 0x00008000
	WNODE_FLAG_PDO_INSTANCE_NAMES    WNODE_HEADER_Flags = 0x00010000
	WNODE_FLAG_TRACED_GUID           WNODE_HEADER_Flags = 0x00020000
	WNODE_FLAG_LOG_WNODE             WNODE_HEADER_Flags = 0x00040000
	WNODE_FLAG_USE_GUID_PTR          WNODE_HEADER_Flags = 0x00080000
	WNODE_FLAG_USE_MOF_PTR           WNODE_HEADER_Flags = 0x00100000
	WNODE_FLAG_NO_HEADER             WNODE_HEADER_Flags = 0x00200000
	WNODE_FLAG_SEND_DATA_BLOCK       WNODE_HEADER_Flags = 0x00400000
	WNODE_FLAG_SEVERITY_MASK         WNODE_HEADER_Flags = 0xff000000
)

// WNODE_HEADER
type WNODE_HEADER struct {
	BufferSize    ULONG
	ProviderId    ULONG
	U1            WNODE_HEADER_u1
	U2            WNODE_HEADER_u2
	Guid          GUID
	ClientContext ULONG
	Flags         WNODE_HEADER_Flags
}

// EventLogFileMode
type EventLogFileMode ULONG

const (
	EVENT_TRACE_FILE_MODE_NONE             EventLogFileMode = 0x00000000
	EVENT_TRACE_FILE_MODE_SEQUENTIAL       EventLogFileMode = 0x00000001
	EVENT_TRACE_FILE_MODE_CIRCULAR         EventLogFileMode = 0x00000002
	EVENT_TRACE_FILE_MODE_APPEND           EventLogFileMode = 0x00000004
	EVENT_TRACE_REAL_TIME_MODE             EventLogFileMode = 0x00000100
	EVENT_TRACE_DELAY_OPEN_FILE_MODE       EventLogFileMode = 0x00000200
	EVENT_TRACE_BUFFERING_MODE             EventLogFileMode = 0x00000400
	EVENT_TRACE_PRIVATE_LOGGER_MODE        EventLogFileMode = 0x00000800
	EVENT_TRACE_ADD_HEADER_MODE            EventLogFileMode = 0x00001000
	EVENT_TRACE_USE_GLOBAL_SEQUENCE        EventLogFileMode = 0x00004000
	EVENT_TRACE_USE_LOCAL_SEQUENCE         EventLogFileMode = 0x00008000
	EVENT_TRACE_RELOG_MODE                 EventLogFileMode = 0x00010000
	EVENT_TRACE_USE_PAGED_MEMORY           EventLogFileMode = 0x01000000
	EVENT_TRACE_FILE_MODE_NEWFILE          EventLogFileMode = 0x00000008
	EVENT_TRACE_FILE_MODE_PREALLOCATE      EventLogFileMode = 0x00000020
	EVENT_TRACE_NONSTOPPABLE_MODE          EventLogFileMode = 0x00000040
	EVENT_TRACE_SECURE_MODE                EventLogFileMode = 0x00000080
	EVENT_TRACE_USE_KBYTES_FOR_SIZE        EventLogFileMode = 0x00002000
	EVENT_TRACE_PRIVATE_IN_PROC            EventLogFileMode = 0x00020000
	EVENT_TRACE_MODE_RESERVED              EventLogFileMode = 0x00100000
	EVENT_TRACE_NO_PER_PROCESSOR_BUFFERING EventLogFileMode = 0x10000000
)

// EventEnableFlags
type EventEnableFlags ULONG

const (
	EVENT_TRACE_FLAG_PROCESS            EventEnableFlags = 0x00000001
	EVENT_TRACE_FLAG_THREAD             EventEnableFlags = 0x00000002
	EVENT_TRACE_FLAG_IMAGE_LOAD         EventEnableFlags = 0x00000004
	EVENT_TRACE_FLAG_DISK_IO            EventEnableFlags = 0x00000100
	EVENT_TRACE_FLAG_DISK_FILE_IO       EventEnableFlags = 0x00000200
	EVENT_TRACE_FLAG_MEMORY_PAGE_FAULTS EventEnableFlags = 0x00001000
	EVENT_TRACE_FLAG_MEMORY_HARD_FAULTS EventEnableFlags = 0x00002000
	EVENT_TRACE_FLAG_NETWORK_TCPIP      EventEnableFlags = 0x00010000
	EVENT_TRACE_FLAG_REGISTRY           EventEnableFlags = 0x00020000
	EVENT_TRACE_FLAG_DBGPRINT           EventEnableFlags = 0x00040000
	EVENT_TRACE_FLAG_PROCESS_COUNTERS   EventEnableFlags = 0x00000008
	EVENT_TRACE_FLAG_CSWITCH            EventEnableFlags = 0x00000010
	EVENT_TRACE_FLAG_DPC                EventEnableFlags = 0x00000020
	EVENT_TRACE_FLAG_INTERRUPT          EventEnableFlags = 0x00000040
	EVENT_TRACE_FLAG_SYSTEMCALL         EventEnableFlags = 0x00000080
	EVENT_TRACE_FLAG_DISK_IO_INIT       EventEnableFlags = 0x00000400
	EVENT_TRACE_FLAG_ALPC               EventEnableFlags = 0x00100000
	EVENT_TRACE_FLAG_SPLIT_IO           EventEnableFlags = 0x00200000
	EVENT_TRACE_FLAG_DRIVER             EventEnableFlags = 0x00800000
	EVENT_TRACE_FLAG_PROFILE            EventEnableFlags = 0x01000000
	EVENT_TRACE_FLAG_FILE_IO            EventEnableFlags = 0x02000000
	EVENT_TRACE_FLAG_FILE_IO_INIT       EventEnableFlags = 0x04000000
	EVENT_TRACE_FLAG_DISPATCHER         EventEnableFlags = 0x00000800
	EVENT_TRACE_FLAG_VIRTUAL_ALLOC      EventEnableFlags = 0x00004000
	EVENT_TRACE_FLAG_EXTENSION          EventEnableFlags = 0x80000000
	EVENT_TRACE_FLAG_FORWARD_WMI        EventEnableFlags = 0x40000000
	EVENT_TRACE_FLAG_ENABLE_RESERVE     EventEnableFlags = 0x20000000
)

// EVENT_TRACE_PROPERTIES
type EVENT_TRACE_PROPERTIES struct {
	Wnode               WNODE_HEADER
	BufferSize          ULONG
	MinimumBuffers      ULONG
	MaximumBuffers      ULONG
	MaximumFileSize     ULONG
	LogFileMode         EventLogFileMode
	FlushTimer          ULONG
	EnableFlags         EventEnableFlags
	AgeLimit            LONG
	NumberOfBuffers     ULONG
	FreeBuffers         ULONG
	EventsLost          ULONG
	BuffersWritten      ULONG
	LogBuffersLost      ULONG
	RealTimeBuffersLost ULONG
	LoggerThreadId      HANDLE
	LogFileNameOffset   ULONG
	LoggerNameOffset    ULONG
}

type PEVENT_TRACE_PROPERTIES *EVENT_TRACE_PROPERTIES

// EVENT_TRACE_HEADER_u1_s
type EVENT_TRACE_HEADER_u1_s struct {
	HeaderType  UCHAR
	MarkerFlags UCHAR
}

// EVENT_TRACE_HEADER_u1
type EVENT_TRACE_HEADER_u1 struct {
	FieldTypeFlags USHORT
}

// EVENT_TRACE_TYPE
type EVENT_TRACE_TYPE UCHAR

const (
	EVENT_TRACE_TYPE_INFO           EVENT_TRACE_TYPE = 0x00
	EVENT_TRACE_TYPE_START          EVENT_TRACE_TYPE = 0x01
	EVENT_TRACE_TYPE_END            EVENT_TRACE_TYPE = 0x02
	EVENT_TRACE_TYPE_STOP           EVENT_TRACE_TYPE = 0x02
	EVENT_TRACE_TYPE_DC_START       EVENT_TRACE_TYPE = 0x03
	EVENT_TRACE_TYPE_DC_END         EVENT_TRACE_TYPE = 0x04
	EVENT_TRACE_TYPE_EXTENSION      EVENT_TRACE_TYPE = 0x05
	EVENT_TRACE_TYPE_REPLY          EVENT_TRACE_TYPE = 0x06
	EVENT_TRACE_TYPE_DEQUEUE        EVENT_TRACE_TYPE = 0x07
	EVENT_TRACE_TYPE_RESUME         EVENT_TRACE_TYPE = 0x07
	EVENT_TRACE_TYPE_CHECKPOINT     EVENT_TRACE_TYPE = 0x08
	EVENT_TRACE_TYPE_SUSPEND        EVENT_TRACE_TYPE = 0x08
	EVENT_TRACE_TYPE_WINEVT_SEND    EVENT_TRACE_TYPE = 0x09
	EVENT_TRACE_TYPE_WINEVT_RECEIVE EVENT_TRACE_TYPE = 0xF0
)

// TRACE_LEVEL
type TRACE_LEVEL UCHAR

const (
	TRACE_LEVEL_NONE        TRACE_LEVEL = 0
	TRACE_LEVEL_FATAL       TRACE_LEVEL = 1
	TRACE_LEVEL_ERROR       TRACE_LEVEL = 2
	TRACE_LEVEL_WARNING     TRACE_LEVEL = 3
	TRACE_LEVEL_INFORMATION TRACE_LEVEL = 4
	TRACE_LEVEL_VERBOSE     TRACE_LEVEL = 5
	TRACE_LEVEL_RESERVED6   TRACE_LEVEL = 6
	TRACE_LEVEL_RESERVED7   TRACE_LEVEL = 7
	TRACE_LEVEL_RESERVED8   TRACE_LEVEL = 8
	TRACE_LEVEL_RESERVED9   TRACE_LEVEL = 9
)

// EVENT_TRACE_HEADER_u2_s
type EVENT_TRACE_HEADER_u2_s struct {
	Type    EVENT_TRACE_TYPE
	Level   TRACE_LEVEL
	Version USHORT
}

// EVENT_TRACE_HEADER_u2
type EVENT_TRACE_HEADER_u2 struct {
	Version ULONG
}

// EVENT_TRACE_HEADER_u3
type EVENT_TRACE_HEADER_u3 struct {
	Guid GUID
}

// EVENT_TRACE_HEADER_u4_s1
type EVENT_TRACE_HEADER_u4_s1 struct {
	KernelTime ULONG
	UserTime   ULONG
}

// EVENT_TRACE_HEADER_u4_s2
type EVENT_TRACE_HEADER_u4_s2 struct {
	ClientContext ULONG
	Flags         WNODE_HEADER_Flags
}

// EVENT_TRACE_HEADER_u4
type EVENT_TRACE_HEADER_u4 struct {
	ProcessorTime ULONG64
}

// EVENT_TRACE_HEADER
type EVENT_TRACE_HEADER struct {
	Size      USHORT
	U1        EVENT_TRACE_HEADER_u1
	U2        EVENT_TRACE_HEADER_u2
	ThreadId  ULONG
	ProcessId ULONG
	TimeStamp LARGE_INTEGER
	U3        EVENT_TRACE_HEADER_u3
	U4        EVENT_TRACE_HEADER_u4
}
type PEVENT_TRACE_HEADER *EVENT_TRACE_HEADER

// EVENT_INSTANCE_HEADER_u1_s
type EVENT_INSTANCE_HEADER_u1_s struct {
	EventId ULONG
	Flags   WNODE_HEADER_Flags
}

// EVENT_INSTANCE_HEADER_u
type EVENT_INSTANCE_HEADER_u struct {
	ProcessorTime ULONG64
}

// EVENT_INSTANCE_HEADER
type EVENT_INSTANCE_HEADER struct {
	Size             USHORT
	U1               EVENT_TRACE_HEADER_u1
	U2               EVENT_TRACE_HEADER_u2
	ThreadId         ULONG
	ProcessId        ULONG
	TimeStamp        LARGE_INTEGER
	RegHandle        ULONGLONG
	InstanceId       ULONG
	ParentInstanceId ULONG
	U                EVENT_INSTANCE_HEADER_u
	ParentRegHandle  ULONGLONG
}
type PEVENT_INSTANCE_HEADER *EVENT_INSTANCE_HEADER

// ProcessTraceMode
type ProcessTraceMode ULONG

const (
	PROCESS_TRACE_MODE_REAL_TIME     ProcessTraceMode = 0x00000100
	PROCESS_TRACE_MODE_RAW_TIMESTAMP ProcessTraceMode = 0x00001000
	PROCESS_TRACE_MODE_EVENT_RECORD  ProcessTraceMode = 0x10000000
)

// EVENT_TRACE_LOGFILE_u1
type EVENT_TRACE_LOGFILE_u1 struct {
	LogFileMode EventLogFileMode
}

// EVENT_TRACE_LOGFILE_u2
type EVENT_TRACE_LOGFILE_u2 struct {
	EventCallback PEVENT_CALLBACK
	// EventRecordCallback PEVENT_RECORD_CALLBACK
}

// ETW_BUFFER_CONTEXT
type ETW_BUFFER_CONTEXT struct {
	ProcessorNumber UCHAR
	Alignment       UCHAR
	LoggerId        USHORT
}

// EVENT_TRACE_u
type EVENT_TRACE_u struct {
	ClientContext ULONG
	// BufferContext ETW_BUFFER_CONTEXT
}

// EVENT_TRACE
type EVENT_TRACE struct {
	Header           EVENT_TRACE_HEADER
	InstanceId       ULONG
	ParentInstanceId ULONG
	ParentGuid       GUID
	MofData          PVOID
	MofLength        ULONG
	U                EVENT_TRACE_u
}

// TRACE_LOGFILE_HEADER_u1_s
type TRACE_LOGFILE_HEADER_u1_s struct {
	MajorVersion    UCHAR
	MinorVersion    UCHAR
	SubVersion      UCHAR
	SubMinorVersion UCHAR
}

// TRACE_LOGFILE_HEADER_u1
type TRACE_LOGFILE_HEADER_u1 struct {
	Version ULONG
	// VersionDetail TRACE_LOGFILE_HEADER_u1_s
}

// TRACE_LOGFILE_HEADER_u2_s
type TRACE_LOGFILE_HEADER_u2_s struct {
	StartBuffers  ULONG
	PointerSize   ULONG
	EventsLost    ULONG
	CpuSpeedInMHz ULONG
}

// TRACE_LOGFILE_HEADER_u2
type TRACE_LOGFILE_HEADER_u2 struct {
	LogInstanceGuid GUID
	//  [TRACE_LOGFILE_HEADER_u2_s]
}

// TRACE_LOGFILE_HEADER
type TRACE_LOGFILE_HEADER struct {
	BufferSize         ULONG
	U1                 TRACE_LOGFILE_HEADER_u1
	ProviderVersion    ULONG
	NumberOfProcessors ULONG
	EndTime            LARGE_INTEGER
	TimerResolution    ULONG
	MaximumFileSize    ULONG
	LogFileMode        EventLogFileMode
	BuffersWritten     ULONG
	U2                 TRACE_LOGFILE_HEADER_u2
	LoggerName         LPWSTR
	LogFileName        LPWSTR
	TimeZone           TIME_ZONE_INFORMATION
	BootTime           LARGE_INTEGER
	PerfFreq           LARGE_INTEGER
	StartTime          LARGE_INTEGER
	ReservedFlags      ULONG
	BuffersLost        ULONG
}

// EVENT_TRACE_LOGFILE
type EVENT_TRACE_LOGFILE struct {
	LogFileName    LPWSTR
	LoggerName     LPWSTR
	CurrentTime    LONGLONG
	BuffersRead    ULONG
	LogFileMode    EventLogFileMode
	CurrentEvent   EVENT_TRACE
	LogfileHeader  TRACE_LOGFILE_HEADER
	BufferCallback PEVENT_TRACE_BUFFER_CALLBACK
	BufferSize     ULONG
	Filled         ULONG
	EventsLost     ULONG
	EventCallback  uintptr
	IsKernelTrace  ULONG
	Context        PVOID
}
type PEVENT_TRACE_LOGFILE *EVENT_TRACE_LOGFILE

// EVENT_INSTANCE_INFO
type EVENT_INSTANCE_INFO struct {
	RegHandle  HANDLE
	InstanceId ULONG
}

type PEVENT_INSTANCE_INFO *EVENT_INSTANCE_INFO

// EVENT_DESCRIPTOR
type EVENT_DESCRIPTOR struct {
	Id      USHORT
	Version UCHAR
	Channel UCHAR
	Level   UCHAR
	Opcode  UCHAR
	Task    USHORT
	Keyword ULONGLONG
}

type PCEVENT_DESCRIPTOR *EVENT_DESCRIPTOR

// EVENT_DATA_DESCRIPTOR
type EVENT_DATA_DESCRIPTOR struct {
	Ptr      ULONGLONG
	Size     ULONG
	Reserved ULONG
}

type PEVENT_DATA_DESCRIPTOR *EVENT_DATA_DESCRIPTOR

// EventLogType
type EventLogType WORD

const (
	EVENTLOG_SUCCESS          EventLogType = 0x0000
	EVENTLOG_ERROR_TYPE       EventLogType = 0x0001
	EVENTLOG_WARNING_TYPE     EventLogType = 0x0002
	EVENTLOG_INFORMATION_TYPE EventLogType = 0x0004
	EVENTLOG_AUDIT_SUCCESS    EventLogType = 0x0008
	EVENTLOG_AUDIT_FAILURE    EventLogType = 0x0010
)

// EventActivity
type EventActivity ULONG

const (
	EVENT_ACTIVITY_CTRL_GET_ID        EventActivity = 1
	EVENT_ACTIVITY_CTRL_SET_ID        EventActivity = 2
	EVENT_ACTIVITY_CTRL_CREATE_ID     EventActivity = 3
	EVENT_ACTIVITY_CTRL_GET_SET_ID    EventActivity = 4
	EVENT_ACTIVITY_CTRL_CREATE_SET_ID EventActivity = 5
)
