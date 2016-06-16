/**
 * Created by igor on 24.05.16.
 */

'use strict';

const MEMBER_STATE = {
    Idle: 0,
    Process: 1,
    End: 2
};

const DIALER_STATES = {
    Idle: 0,
    Work: 1,
    Sleep: 2,
    ProcessStop: 3,
    End: 4,
    Error: 5
};

const DIALER_CAUSE = {
    Init: "INIT",
    ProcessStop: "QUEUE_STOP",
    ProcessRecovery: "QUEUE_RECOVERY",
    ProcessSleep: "QUEUE_SLEEP",
    ProcessReady: "QUEUE_HUNTING",
    ProcessNotFoundMember: "NOT_FOUND_MEMBER",
    ProcessComplete: "QUEUE_COMPLETE",
    ProcessExpire: "QUEUE_EXPIRE",
    ProcessInternalError: "QUEUE_ERROR"
};

const DIFF_CHANGE_MSEC = 2000;

const AGENT_STATUS = {
    LoggedOut: 'Logged Out',
    Available: 'Available',
    AvailableOnDemand: 'Available (On Demand)',
    OnBreak: 'On Break'
};

const AGENT_STATE = {
    // Idle: 'Reserved',
    Idle: 'Idle',
    Reserved: 'Reserved',
    Waiting: 'Waiting',
    InQueueCall: 'In a queue call'
};

const END_CAUSE = {
    NO_ROUTE: "NO_ROUTE",
    MAX_TRY: "MAX_TRY_COUNT",
    PROCESS_CRASH: "PROCESS_CRASH",
    ACCEPT: "ACCEPT"
};

const CODE_RESPONSE_ERRORS = ["UNALLOCATED_NUMBER", "NORMAL_TEMPORARY_FAILURE", END_CAUSE.NO_ROUTE, 'CHAN_NOT_IMPLEMENTED', "CALL_REJECTED", "INVALID_NUMBER_FORMAT", "NETWORK_OUT_OF_ORDER", "NORMAL_TEMPORARY_FAILURE", "OUTGOING_CALL_BARRED", "SERVICE_UNAVAILABLE", "CHAN_NOT_IMPLEMENTED", "SERVICE_NOT_IMPLEMENTED", "INCOMPATIBLE_DESTINATION", "MANDATORY_IE_MISSING", "PROGRESS_TIMEOUT", "GATEWAY_DOWN"];
const CODE_RESPONSE_RETRY = ["NO_ROUTE_DESTINATION", "DESTINATION_OUT_OF_ORDER", "USER_BUSY", "NO_USER_RESPONSE", "NO_ANSWER", "SUBSCRIBER_ABSENT", "NUMBER_CHANGED", "NORMAL_UNSPECIFIED", "NORMAL_CIRCUIT_CONGESTION", "ORIGINATOR_CANCEL", "LOSE_RACE", "USER_NOT_REGISTERED"];
const CODE_RESPONSE_OK = ["NORMAL_CLEARING"];

const MAX_MEMBER_RETRY = 999;

const DIALER_TYPES = {
    VoiceBroadcasting: "Voice Broadcasting",
    ProgressiveDialer: "Progressive Dialer",
    PredictiveDialer: "Predictive Dialer"
};

module.exports = {
    DIALER_STATES: DIALER_STATES,
    DIALER_CAUSE: DIALER_CAUSE,
    DIFF_CHANGE_MSEC: DIFF_CHANGE_MSEC,
    AGENT_STATUS: AGENT_STATUS,
    AGENT_STATE: AGENT_STATE,
    END_CAUSE: END_CAUSE,
    CODE_RESPONSE_ERRORS: CODE_RESPONSE_ERRORS,
    CODE_RESPONSE_RETRY: CODE_RESPONSE_RETRY,
    CODE_RESPONSE_OK: CODE_RESPONSE_OK,
    MAX_MEMBER_RETRY: MAX_MEMBER_RETRY,
    DIALER_TYPES: DIALER_TYPES,
    MEMBER_STATE: MEMBER_STATE
};