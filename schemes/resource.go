package schemes

type Permission = string
type MetaFileType = string

const (
	ScreenCapture   Permission = "Screen Capture"
	WebRTC          Permission = "WebRTC"
	ClipboardAccess Permission = "Clipboard Access"
)

const (
	Carcols         MetaFileType = "CARCOLS_FILE"
	VehicleMetaData MetaFileType = "VEHICLE_METADATA_FILE"
	Handling        MetaFileType = "HANDLING_FILE"
)

// ResourceConfig resource configuration - resource.cfg
type ResourceConfig struct {
	// Type The serverside type of your resource (the correct module for that type has to be loaded)
	Type string `alt:"type"`
	// ClientType The clientside type of your resource (the correct module for that type has to be loaded)
	ClientType string `alt:"client-type"`
	// Main The main serverside file that will get loaded when the server starts
	Main string `alt:"main"`
	// ClientMain The main clientside file that will get loaded when the client starts
	ClientMain string `alt:"client-main"`
	// ClientFiles The files that the client has access to (The client-main file does not have to be included here)
	ClientFiles []string `alt:"client-files"`
	// RequiredPermissions The required permissions to play on the server (these have to be accepted, otherwise the player can't join)
	RequiredPermissions []Permission `alt:"required-permissions"`
	// OptionalPermissions The optional permissions to play on the server (these permissions can be declined by the user)
	OptionalPermissions []Permission `alt:"optional-permissions"`
	// Dependencies The dependencies of this resource (All dependencies get loaded before the resource)
	Dependencies []string `alt:"deps"`
}

// StreamConfig dlc resource configuration - stream.cfg
type StreamConfig struct {
	// An array containing all Files that should be loaded for this dlc
	// You can also use glob patterns to give access to a whole directory
	// Example: "stream/assets/*"
	Files []string `alt:"files"`
	// The Meta files your dlc uses
	Meta map[string]MetaFileType `alt:"meta"`
	// The GXT entries for your dlc
	GXT map[string]string `alt:"gxt"`
}
