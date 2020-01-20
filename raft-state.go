package raftState

type common struct {
        /*
         * Persistent State on all Servers
         */
        currentTerm int32    /* latest term server has seen (initialized to 0
                              * on first boot, increases monotonically) */
        votedFor    int32    /* candidateId that received vote in current
                              * term (or null if none) */
        raftLog     []string /* log entries; each entry contains command
                              * for state machine, and term when entry
                              * was received by leader (first index is 1) */
        /*
         * Volatile State on all Servers
         * (Reinitialized after election)
         */
        commitIndex int32 /* index of highest log entry known to be
                           * committed (initialized to 0, increases
                           * monotonically) */
        lastApplied int32 /* index of highest log entry applied to state
                           * machine (initialized to 0, increases
                           * monotonically) */
}

type leaderState []struct {
        /*
         * Volatile state on leaders
         */
        nextIndex  int32 /* for each server, index of the next log entry
                          * to send to that server (initialized to leader
                          * last log index + 1) */
        matchIndex int32 /* or each server, index of highest log entry
                          * known to be replicated on server
                          * (initialized to 0, increases monotonically) */
}

type Actions int32
const (
        timer Actions = iota
        message
)

const (
        port = ":34567"
)
