The message serialization involves an outer JSON envelope rooted in a map.  The base protocol only defines a single field in that map, "v" which represents the protocol version as an ordinal number.  Each version number is controlling and opinionated about the remaining field names, field types, signature algorithms, encryption algorithms, including source and destination address formats.  Implementations must handle at least INT32_MAX version numbers.  Based on the version of the message protocol, the payload will be embedded in a protocol defined manner in the envelope JSON.

Implementations will silently discard any messages with a version number they don't support / understand.

Protocol version 0 is called "Clarity" and its purpose is to have verbose JSON for ease of development and debugging.  Payloads are inline and not encrypted, encoded, or compressed.  Checksum is a basic CRC-32 rather than a signature/tag/hash.  Protocol v0 relies on SRV records to bootstrap, optionally implementing host discovery after bootstrapping.  Protocol v0 also relies on the well-known URI path "/.mesh/v0" which is used when forming URIs from the SRV records.

Protocol version 1 will be called "Puny" and is expected to be minimized JSON (request/response IDs are UUIDv4 encoded in Base58, field names are single or double letters, use Base122 encoding rather than Base64 for encrypted payload, timestamps in Unix timestamp with fractional seconds. The payload JSON will be compressed with gzip-1 / fastest before encryption.


v0 serialization format (Unencrypted inline payload)
{
	v: "0",									# Protocol version 0 ("Clarity") - Nodes discard versions they don't understand
	TTL: "4",								# TTL { int8 } (Decremented each time the message is received, if zero, message is discarded)
	SourceId: "b8ae839475",					# Source { Opaque string } (Randomly generated node identifier)
	DestinationId: "a732bcd89b",			# Destination { Opaque string } (Randomly generated node identifier, "*" for broadcast, until TTL is 0)
	Checksum: "0xCBF43926"					# Checksum of [ Source, Destination, Payload ] fields only { CRC32 }
	InlinePayload: {						# For protocol version 0, the payload is inline JSON and not encrypted
		RequestID: "3",						# Request identifier for response { Opaque string, in64 sequence in v0 }
		MessageType: "ping",
		Headers: {
			SentTimestamp: "2021-05-09T17:28:07-07:00"
			SendPath: "https://myhost.com:443/.mesh/v0"
		}
	}
}

Payload formats

{
	RequestID: "3"										# Request identifier { Opaque string, int64 sequence in v0 }
	MessageType: "ping"									# Message Type { Enum: ping, pong, announce }
	Headers: {											# Headers { map }
		SentTimestamp: "2021-05-09T17:28:07-07:00"		# { Localized ISO 8601 with offset }
		SendPath: "https://myhost.com:443/.mesh/v0"		# The path the sender used to send this ping
	}
}

{
	ResponseID: "3"												# Responding to request identifier { Opaque string, int64 sequence in v0 }
	MessageType: "pong"											# Message Type { Enum: ping, pong, announce }
	Headers: {													# Headers { map }
		SentTimestamp: "2021-05-09T17:28:07.38-07:00",			# Sent timestamp received in Ping message { Localized ISO 8601 with offset }
		ReceivedTimestamp: "2021-05-09T17:28:07.98-07:00"		# Received timestamp for Ping message { Localized ISO 8601 with offset }
		SendPath: "https://74.123.80.80:4443/.mesh/v0"		# The path the sender used to send this ping
	}
}

{
	MessageType: "announce"
	Headers: {
		SentTimestamp: "2020-12-12T19:38:02.34-07:00"			# { Localized ISO 8601 with offset }
		SendPath: "https://myhost.com:443/.mesh/v0"				# The path the sender used to send this announcement
	}
	Paths: [													# List of possible paths to this node
			{
				Path: "https://74.123.80.80:4443/.mesh/v0",
				Priority: "1",									# Priority of the target host, lower value means more preferred
				Weight: "100"									# Relative weight for records of same priority, higher value, higher chance of getting picked
			},
			{
				Path: "wss://74.123.80.80:4443/.mesh/v0",
				Priority: "2",
				Weight: "100"
			}
	],
	Peers: [													# List of known peers and their paths
		{
			Id: "b8ae839475",
			Paths [
				{
					Path:  "https://94.127.8.39/.mesh/v0"
					Priority: "1"
					Weight: "100"
				}
			]
		}
	]
}







Future announce might look like:


{
	MessageType: "announce"
	Headers: {
		SentTimestamp: "2020-12-12T19:38:02.34-07:00"			# { Localized ISO 8601 with offset }
	}
	Capabilities: {												# Node advertizes capabilities it is able to provide other nodes
		SendEmail: {											# Node advertizes that it is configured so it can send emails
			FromAddress: "mesh@mydomain.com",					# Source e-mail address the Node emails will come from
			MaxSize: "10485760"									# Maximum email size in bytes
		}
		Dns: {											# Node advertises that it is able to create / configure DNS records
			Subdomain: "*.mydomain.com"					# Node advertises that it can create arbitrary subdomains under this domain
		}
	}
	Traits: {
		PowerSource: "battery"							# Node advertises that it is on battery
		CpuArchitecture: "aarch64"						# Node advertises ARM 64bit CPU
		Memory:	"17179869184"							# Node advertises 16GiB memory
		OSFamily: "Linux"								# Node advertises [ Windows, MacOS, Linux ]
	}
	Paths: [
			{
				Path: "https://74.123.80.80:4443/.mesh/v0",
				Priority: "1",
				Weight: "100"
			},
			{
				Path: "wss://74.123.80.80:4443/.mesh/v0",
				Priority: "2",
				Weight: "100"
			}
	],
	Peers: [
		{
			Id: "b8ae839475",
			Paths [
				{
					Path:  "https://94.127.8.39/.mesh/v0"
					Priority: "1"
					Weight: "100"
				}
			]
		}
	]
}
