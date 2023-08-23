import { isEnableDebug } from './settings.js';

function logDebug(msg, obj) {
    if (isEnableDebug()) {
        if (obj) {
            console.debug('[n.f97.xyz Debug] ' + msg, obj);
        } else {
            console.debug('[n.f97.xyz Debug] ' + msg);
        }
    }
}

function logInfo(msg, obj) {
    if (obj) {
        console.info('[n.f97.xyz Info] ' + msg, obj);
    } else {
        console.info('[n.f97.xyz Info] ' + msg);
    }
}

function logWarn(msg, obj) {
    if (obj) {
        console.warn('[n.f97.xyz Warn] ' + msg, obj);
    } else {
        console.warn('[n.f97.xyz Warn] ' + msg);
    }
}

function logError(msg, obj) {
    if (obj) {
        console.error('[n.f97.xyz Error] ' + msg, obj);
    } else {
        console.error('[n.f97.xyz Error] ' + msg);
    }
}

export default {
    debug: logDebug,
    info: logInfo,
    warn: logWarn,
    error: logError
};
