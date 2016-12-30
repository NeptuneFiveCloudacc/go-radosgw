package radosAPI

type apiError struct {
	Code string `json:"Code"`
}

// Usage represents the response of usage requests
type Usage struct {
	Entries []struct {
		Buckets []struct {
			Bucket     string `json:"bucket"`
			Categories []struct {
				BytesReceived int    `json:"bytes_received"`
				BytesSent     int    `json:"bytes_sent"`
				Category      string `json:"category"`
				Ops           int    `json:"ops"`
				SuccessfulOps int    `json:"successful_ops"`
			} `json:"categories"`
			Epoch int    `json:"epoch"`
			Time  string `json:"time"`
		} `json:"buckets"`
		Owner string `json:"owner"`
	} `json:"entries"`
	Summary []struct {
		Categories []struct {
			BytesReceived int    `json:"bytes_received"`
			BytesSent     int    `json:"bytes_sent"`
			Category      string `json:"category"`
			Ops           int    `json:"ops"`
			SuccessfulOps int    `json:"successful_ops"`
		} `json:"categories"`
		Total struct {
			BytesReceived int `json:"bytes_received"`
			BytesSent     int `json:"bytes_sent"`
			Ops           int `json:"ops"`
			SuccessfulOps int `json:"successful_ops"`
		} `json:"total"`
		User string `json:"user"`
	} `json:"summary"`
}

// SubUsers represents the response of subuser requests
type SubUsers []struct {
	ID          string `json:"id"`
	Permissions string `json:"permissions"`
}

// KeysDefinition represents the response of key requests
type KeysDefinition []struct {
	AccessKey string `json:"access_key,omitempty"`
	SecretKey string `json:"secret_key"`
	User      string `json:"user"`
}

// User represents the response of user requests
type User struct {
	Caps []struct {
		Perm string `json:"perm"`
		Type string `json:"type"`
	} `json:"caps"`
	DisplayName string         `json:"display_name"`
	Email       string         `json:"email"`
	Keys        KeysDefinition `json:"keys"`
	MaxBuckets  int            `json:"max_buckets"`
	Subusers    SubUsers       `json:"subusers"`
	Suspended   int            `json:"suspended"`
	SwiftKeys   KeysDefinition `json:"swift_keys"`
	UserID      string         `json:"user_id"`
}

type stats struct {
	Bucket      string `json:"bucket"`
	BucketQuota struct {
		Enabled    bool `json:"enabled"`
		MaxObjects float64  `json:"max_objects"`
		MaxSizeKb  float64  `json:"max_size_kb"`
	} `json:"bucket_quota"`
	ID        string `json:"id"`
	IndexPool string `json:"index_pool"`
	Marker    string `json:"marker"`
	MasterVer string `json:"master_ver"`
	MaxMarker string `json:"max_marker"`
	Mtime     string `json:"mtime"`
	Owner     string `json:"owner"`
	Pool      string `json:"pool"`
	Usage     struct {
		RgwMain struct {
			NumObjects   float64 `json:"num_objects"`
			SizeKb       float64 `json:"size_kb"`
			SizeKbActual float64 `json:"size_kb_actual"`
		} `json:"rgw.main"`
	} `json:"usage"`
	Ver string `json:"ver"`
}

type bucket struct {
	Name  string `json:"name,omitempty"`
	Stats *stats `json:"stats,omitempty"`
}

// Buckets represents the response of bucket requests
type Buckets []bucket
