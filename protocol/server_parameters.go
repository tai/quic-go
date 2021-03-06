package protocol

import "time"

// DefaultMaxCongestionWindow is the default for the max congestion window
const DefaultMaxCongestionWindow = 1000

// InitialCongestionWindow is the initial congestion window in QUIC packets
const InitialCongestionWindow = 32

// MaxUndecryptablePackets limits the number of undecryptable packets that a
// session queues for later until it sends a public reset.
const MaxUndecryptablePackets = 10

// AckSendDelay is the maximum delay that can be applied to an ACK for a retransmittable packet
// This is the value Chromium is using
const AckSendDelay = 25 * time.Millisecond

// ReceiveStreamFlowControlWindow is the stream-level flow control window for receiving data
// This is the value that Google servers are using
const ReceiveStreamFlowControlWindow ByteCount = (1 << 10) * 32 // 32 kB

// ReceiveConnectionFlowControlWindow is the connection-level flow control window for receiving data
// This is the value that Google servers are using
const ReceiveConnectionFlowControlWindow ByteCount = (1 << 10) * 48 // 48 kB

// MaxReceiveStreamFlowControlWindow is the maximum stream-level flow control window for receiving data
// This is the value that Google servers are using
const MaxReceiveStreamFlowControlWindow ByteCount = 1 * (1 << 20) // 1 MB

// MaxReceiveConnectionFlowControlWindow is the connection-level flow control window for receiving data
// This is the value that Google servers are using
const MaxReceiveConnectionFlowControlWindow ByteCount = 1.5 * (1 << 20) // 1.5 MB

// MaxStreamsPerConnection is the maximum value accepted for the number of streams per connection
const MaxStreamsPerConnection = 100

// MaxIncomingDynamicStreamsPerConnection is the maximum value accepted for the incoming number of dynamic streams per connection
const MaxIncomingDynamicStreamsPerConnection = 100

// MaxStreamsMultiplier is the slack the client is allowed for the maximum number of streams per connection, needed e.g. when packets are out of order or dropped. The minimum of this procentual increase and the absolute increment specified by MaxStreamsMinimumIncrement is used.
const MaxStreamsMultiplier = 1.1

// MaxStreamsMinimumIncrement is the slack the client is allowed for the maximum number of streams per connection, needed e.g. when packets are out of order or dropped. The minimum of this absolute increment and the procentual increase specified by MaxStreamsMultiplier is used.
const MaxStreamsMinimumIncrement = 10

// MaxNewStreamIDDelta is the maximum difference between and a newly opened Stream and the highest StreamID that a client has ever opened
// note that the number of streams is half this value, since the client can only open streams with open StreamID
const MaxNewStreamIDDelta = 4 * MaxStreamsPerConnection

// MaxSessionUnprocessedPackets is the max number of packets stored in each session that are not yet processed.
const MaxSessionUnprocessedPackets = DefaultMaxCongestionWindow

// RetransmissionThreshold + 1 is the number of times a packet has to be NACKed so that it gets retransmitted
const RetransmissionThreshold = 3

// SkipPacketAveragePeriodLength is the average period length in which one packet number is skipped to prevent an Optimistic ACK attack
const SkipPacketAveragePeriodLength PacketNumber = 500

// MaxTrackedSkippedPackets is the maximum number of skipped packet numbers the SentPacketHandler keep track of for Optimistic ACK attack mitigation
const MaxTrackedSkippedPackets = 10

// STKExpiryTimeSec is the valid time of a source address token in seconds
const STKExpiryTimeSec = 24 * 60 * 60

// MaxTrackedSentPackets is maximum number of sent packets saved for either later retransmission or entropy calculation
const MaxTrackedSentPackets = 2 * DefaultMaxCongestionWindow

// MaxTrackedReceivedPackets is the maximum number of received packets saved for doing the entropy calculations
const MaxTrackedReceivedPackets = 2 * DefaultMaxCongestionWindow

// MaxTrackedReceivedAckRanges is the maximum number of ACK ranges tracked
const MaxTrackedReceivedAckRanges = DefaultMaxCongestionWindow

// MaxPacketsReceivedBeforeAckSend is the number of packets that can be received before an ACK frame is sent
const MaxPacketsReceivedBeforeAckSend = 20

// RetransmittablePacketsBeforeAck is the number of retransmittable that an ACK is sent for
const RetransmittablePacketsBeforeAck = 2

// MaxStreamFrameSorterGaps is the maximum number of gaps between received StreamFrames
// prevents DoS attacks against the streamFrameSorter
const MaxStreamFrameSorterGaps = 1000

// CryptoMaxParams is the upper limit for the number of parameters in a crypto message.
// Value taken from Chrome.
const CryptoMaxParams = 128

// CryptoParameterMaxLength is the upper limit for the length of a parameter in a crypto message.
const CryptoParameterMaxLength = 2000

// EphermalKeyLifetime is the lifetime of the ephermal key during the handshake, see handshake.getEphermalKEX.
const EphermalKeyLifetime = time.Minute

// InitialIdleTimeout is the timeout before the handshake succeeds.
const InitialIdleTimeout = 5 * time.Second

// DefaultIdleTimeout is the default idle timeout.
const DefaultIdleTimeout = 30 * time.Second

// MaxIdleTimeout is the maximum idle timeout that can be negotiated.
const MaxIdleTimeout = 1 * time.Minute

// MaxTimeForCryptoHandshake is the default timeout for a connection until the crypto handshake succeeds.
const MaxTimeForCryptoHandshake = 10 * time.Second

// ClosedSessionDeleteTimeout the server ignores packets arriving on a connection that is already closed
// after this time all information about the old connection will be deleted
const ClosedSessionDeleteTimeout = time.Minute

// NumCachedCertificates is the number of cached compressed certificate chains, each taking ~1K space
const NumCachedCertificates = 128
