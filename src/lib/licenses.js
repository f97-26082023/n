export default {
    getLicense: () => {
        return __GOFIRE_LICENSE__; // eslint-disable-line
    },
    getThirdPartyLicenses: () => {
        return __GOFIRE_THIRD_PARTY_LICENSES__ || []; // eslint-disable-line
    }
};
