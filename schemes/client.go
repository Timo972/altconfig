package schemes

// ClientConfig altv.exe configuration - altv.cfg
type ClientConfig struct {
	// Name Your name that is displayed on a server.
	Name string `alt:"name" default:"alt:V nickname"`
	// Branch The branch on which the client will work.
	Branch string `alt:"branch" default:"release"`
	// Build Client will use this build of the declared branch (auto generated by the client normally).
	Build int `alt:"build" default:"789"`
	// Debug Activates the debug mode. For example, a active debug mode enables the debug-console (accessible via F8) and allows you to use the reconnect command in the console.
	Debug bool `alt:"debug" default:"false"`
	// GTAPath Path to your GTA5 directory. Usually, it is set up automatically thought the installation process.
	GTAPath string `alt:"gtapath"`
	// Lang Language of your client.
	Lang string `alt:"lang" default:"en"`
	// LastIP The ip of the last server you played on.
	LastIP string `alt:"lastip" default:"0.0.0.0"`
	// NetGraph Shows a netgraph on the bottom left.
	NetGraph bool `alt:"netgraph" default:"false"`
	// StreamerMode Enables or disables the streamer mode.
	StreamerMode bool `alt:"streamerMode" default:"false"`
	// UseExternalConsole Enables or disables the external console (console popout).
	UseExternalConsole bool `alt:"useExternalConsole" default:"false"`
	// VoiceActivationKey Sets the key for Push-to-talk. You can get the key code here.
	VoiceActivationKey int `alt:"voiceActivationKey" default:"78"`
	// VoiceActivationEnabled Enables or disables the voice activity input mode.
	VoiceActivationEnabled bool `alt:"voiceActivationEnabled" default:"false"`
	// VoiceInputSensitivity If voiceActivationEnabled is enabled, this option will set the required sensitivity.
	VoiceInputSensitivity int `alt:"voiceInputSensitivity" default:"20"`
	// VoiceEnabled Enables or disables the voice system for the client.
	VoiceEnabled bool `alt:"voiceEnabled" default:"true"`
	// VoiceAutoInputVolume Enables or disables the automatic determination of the input volume.
	VoiceAutoInputVolume bool `alt:"voiceAutoInputVolume" default:"true"`
	// VoiceInputVolume Sets the input volume (Range: 0 - 200).
	VoiceInputVolume uint8 `alt:"voiceInputVolume" default:"100"`
	// VoiceNoiseSuppression Enables or disables the noise suppression.
	VoiceNoiseSuppression bool `alt:"voiceNoiseSuppression" default:"true"`
	// VoiceVolume Sets the output volume (Range: 0 - 200).
	VoiceVolume uint8 `alt:"voiceVolume" default:"100"`
	// EarlyAuthTestURL URL to your early auth website. Only usable in rc & dev branch.
	EarlyAuthTestURL string `alt:"earlyAuthTestURL"`
}
