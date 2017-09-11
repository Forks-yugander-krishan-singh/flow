package protocol

const (
	HeaderDatumType    = "Fnproject-Datumtype"
	HeaderResultStatus = "Fnproject-Resultstatus"
	HeaderResultCode   = "Fnproject-Resultcode"
	HeaderStageRef     = "Fnproject-Stageid"
	HeaderHookRef      = "Fnproject-Hookid"
	HeaderMethod       = "Fnproject-Method"
	HeaderHeaderPrefix = "Fnproject-Header-"
	HeaderErrorType    = "Fnproject-Errortype"
	HeaderStateType    = "Fnproject-Statetype"
	HeaderThreadId     = "Fnproject-Threadid"
	HeaderCodeLocation = "Fnproject-Codeloc"

	HeaderContentType = "Content-Type"

	ResultStatusSuccess = "success"
	ResultStatusFailure = "failure"

	DatumTypeBlob     = "blob"
	DatumTypeEmpty    = "empty"
	DatumTypeError    = "error"
	DatumTypeStageRef = "stageref"
	DatumTypeHttpReq  = "httpreq"
	DatumTypeHttpResp = "httpresp"
	DatumTypeState    = "state"
)
