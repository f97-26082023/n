import { isEnableDebug } from './settings.js';

function logDebug(msg, obj) {
    if (isEnableDebug()) {
        if (obj) {
            console.debug('[mn.f97.xyz Debug] ' + msg, obj);
        } else {
            console.debug('[mn.f97.xyz Debug] ' + msg);
        }
    }
}

function logInfo(msg, obj) {
    if (obj) {
        console.info('[mn.f97.xyz Info] ' + msg, obj);
    } else {
        console.info('[mn.f97.xyz Info] ' + msg);
    }
}

function logWarn(msg, obj) {
    if (obj) {
        console.warn('[mn.f97.xyz Warn] ' + msg, obj);
    } else {
        console.warn('[mn.f97.xyz Warn] ' + msg);
    }
}

function logError(msg, obj) {
    if (obj) {
        console.error('[mn.f97.xyz Error] ' + msg, obj);
    } else {
        console.error('[mn.f97.xyz Error] ' + msg);
    }
}

export default {
    debug: logDebug,
    info: logInfo,
    warn: logWarn,
    error: logError
};
