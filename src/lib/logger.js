import { isEnableDebug } from './settings.js';

function logDebug(msg, obj) {
    if (isEnableDebug()) {
        if (obj) {
            console.debug('[gofire Debug] ' + msg, obj);
        } else {
            console.debug('[gofire Debug] ' + msg);
        }
    }
}

function logInfo(msg, obj) {
    if (obj) {
        console.info('[gofire Info] ' + msg, obj);
    } else {
        console.info('[gofire Info] ' + msg);
    }
}

function logWarn(msg, obj) {
    if (obj) {
        console.warn('[gofire Warn] ' + msg, obj);
    } else {
        console.warn('[gofire Warn] ' + msg);
    }
}

function logError(msg, obj) {
    if (obj) {
        console.error('[gofire Error] ' + msg, obj);
    } else {
        console.error('[gofire Error] ' + msg);
    }
}

export default {
    debug: logDebug,
    info: logInfo,
    warn: logWarn,
    error: logError
};
