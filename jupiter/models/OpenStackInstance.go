package models

type OpenstackInstace struct {

	// TenantID identifies the tenant owning this server resource.
	TenantID string `mapstructure:"tenant_id"`

	// UserID uniquely identifies the user account owning the tenant.
	UserID string `mapstructure:"user_id"`

	// Name contains the human-readable name for the server.
	Name string

	// Updated and Created contain ISO-8601 timestamps of when the state of the server last changed, and when it was created.
	Updated string
	Created string

	HostID string



	// Progress ranges from 0..100.
	// A request made against the server completes only once Progress reaches 100.
	Progress int

	// AccessIPv4 and AccessIPv6 contain the IP addresses of the server, suitable for remote access for administration.
	AccessIPv4, AccessIPv6 string

	// Image refers to a JSON object, which itself indicates the OS image used to deploy the server.
	//Image map[string]interface{}

	// Flavor refers to a JSON object, which itself indicates the hardware configuration of the deployed server.
	//Flavor map[string]interface{}

	// Addresses includes a list of all IP addresses assigned to the server, keyed by pool.
	//Addresses map[string]interface{}

	// Metadata includes a list of all user-specified key-value pairs attached to the server.
	//Metadata map[string]interface{}

	// Links includes HTTP references to the itself, useful for passing along to other APIs that might want a server reference.
	//Links []interface{}

	// KeyName indicates which public key was injected into the server on launch.
	//KeyName string `json:"key_name" mapstructure:"key_name"`

	// AdminPass will generally be empty ("").  However, it will contain the administrative password chosen when provisioning a new server without a set AdminPass setting in the first place.
	// Note that this is the ONLY time this field will be valid.
	//AdminPass string `json:"adminPass" mapstructure:"adminPass"`

	// SecurityGroups includes the security groups that this instance has applied to it
	//SecurityGroups []map[string]interface{} `json:"security_groups" mapstructure:"security_groups"`
}
