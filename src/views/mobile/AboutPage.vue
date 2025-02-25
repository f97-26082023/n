<template>
    <f7-page>
        <f7-navbar :title="$t('About')" :back-link="$t('Back')"></f7-navbar>

        <f7-list strong inset dividers class="margin-top">
            <f7-list-item :title="$t('Version')" :after="version"></f7-list-item>
            <f7-list-item :title="$t('Build Time')" :after="buildTime" v-if="buildTime"></f7-list-item>
            <f7-list-item external :title="$t('Official Website')" link="https://github.com/f97/gofire" target="_blank"></f7-list-item>
            <f7-list-item :title="$t('License')" link="#" popup-open=".license-popup"></f7-list-item>
        </f7-list>

        <f7-popup push with-subnavbar swipe-to-close swipe-handler=".swipe-handler" class="license-popup">
            <f7-page>
                <f7-navbar>
                    <div class="swipe-handler" style="z-index: 10"></div>
                    <f7-subnavbar :title="$t('License') "></f7-subnavbar>
                </f7-navbar>
                <f7-block strong outline class="license-content">
                    <p>
                        <span :key="num" v-for="(line, num) in licenseLines"
                              :style="{ 'display': line ? 'initial' : 'block', 'padding' : line ? '0' : '0 0 1em 0' }">
                            {{ line }}
                        </span>
                    </p>
                    <hr/>
                    <p>
                        <span>gofire also contains additional third party software and illustration.</span><br/>
                        <span>All the third party software/illustration included or linked is redistributed under the terms and conditions of their original licenses.</span>
                    </p>
                    <p></p>
                    <p :key="license.name" v-for="license in thirdPartyLicenses">
                        <strong>{{ license.name }}</strong>
                        <br v-if="license.copyright"/><span v-if="license.copyright">{{ license.copyright }}</span>
                        <br v-if="license.url"/><span class="work-break-all" v-if="license.url">{{ license.url }}</span>
                        <br v-if="license.licenseUrl"/><span class="work-break-all" v-if="license.licenseUrl">License: {{ license.licenseUrl }}</span>
                    </p>
                </f7-block>
            </f7-page>
        </f7-popup>
    </f7-page>
</template>

<script>
import { mapStores } from 'pinia';
import { useUserStore } from '@/stores/user.js';

import licenses from '@/lib/licenses.js';

export default {
    computed: {
        ...mapStores(useUserStore),
        version() {
            return 'v' + this.$version;
        },
        buildTime() {
            if (!this.$buildTime) {
                return this.$buildTime;
            }

            return this.$locale.formatUnixTimeToLongDateTime(this.userStore, this.$buildTime);
        },
        licenseLines() {
            return licenses.getLicense().replaceAll(/\r/g, '').split('\n');
        },
        thirdPartyLicenses() {
            return licenses.getThirdPartyLicenses();
        }
    }
}
</script>

<style>
.license-popup .navbar-bg {
    background-color: rgb(var(--f7-navbar-bg-color-rgb, var(--f7-bars-bg-color-rgb)));
}

.license-popup .subnavbar {
    background-color: rgb(var(--f7-subnavbar-bg-color-rgb, var(--f7-bars-bg-color-rgb)));
}

.license-popup .subnavbar-title {
    --f7-subnavbar-title-font-size: var(--ebk-license-popup-title-font-size);
}

.license-content {
    font-size: var(--ebk-license-content-font-size);
}
</style>
