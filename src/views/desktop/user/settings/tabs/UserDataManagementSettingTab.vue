<template>
    <v-row>
        <v-col cols="12">
            <v-card :class="{ 'disabled': loadingDataStatistics }">
                <template #title>
                    <div class="d-flex align-center">
                        <span>{{ $t('Data Management') }}</span>
                        <v-btn density="compact" color="default" variant="text"
                               class="ml-2" :icon="true"
                               v-if="!loadingDataStatistics" @click="reloadUserDataStatistics(true)">
                            <v-icon :icon="icons.refresh" size="24" />
                            <v-tooltip activator="parent">{{ $t('Refresh') }}</v-tooltip>
                        </v-btn>
                        <v-progress-circular indeterminate size="24" class="ml-2" v-if="loadingDataStatistics"></v-progress-circular>
                    </div>
                </template>

                <v-card-text>
                    <v-row>
                        <v-col cols="6" sm="3">
                            <div class="d-flex align-center">
                                <div class="me-3">
                                    <v-avatar rounded color="info" size="42" class="elevation-1">
                                        <v-icon size="24" :icon="icons.transactions"/>
                                    </v-avatar>
                                </div>

                                <div class="d-flex flex-column">
                                    <span class="text-caption">{{ $t('Transaction') }}</span>
                                    <v-skeleton-loader class="skeleton-no-margin pt-3 pb-2" type="text" style="width: 60px" :loading="true" v-if="loadingDataStatistics"></v-skeleton-loader>
                                    <span class="text-h6" v-if="!loadingDataStatistics">{{ displayDataStatistics ? displayDataStatistics.totalTransactionCount : '-' }}</span>
                                </div>
                            </div>
                        </v-col>
                        <v-col cols="6" sm="3">
                            <div class="d-flex align-center">
                                <div class="me-3">
                                    <v-avatar rounded color="primary" size="42" class="elevation-1">
                                        <v-icon size="24" :icon="icons.accounts"/>
                                    </v-avatar>
                                </div>

                                <div class="d-flex flex-column">
                                    <span class="text-caption">{{ $t('Accounts') }}</span>
                                    <v-skeleton-loader class="skeleton-no-margin pt-3 pb-2" type="text" style="width: 60px" :loading="true" v-if="loadingDataStatistics"></v-skeleton-loader>
                                    <span class="text-h6" v-if="!loadingDataStatistics">{{ displayDataStatistics ? displayDataStatistics.totalAccountCount : '-' }}</span>
                                </div>
                            </div>
                        </v-col>
                        <v-col cols="6" sm="3">
                            <div class="d-flex align-center">
                                <div class="me-3">
                                    <v-avatar rounded color="success" size="42" class="elevation-1">
                                        <v-icon size="24" :icon="icons.categories"/>
                                    </v-avatar>
                                </div>

                                <div class="d-flex flex-column">
                                    <span class="text-caption">{{ $t('Transaction Categories') }}</span>
                                    <v-skeleton-loader class="skeleton-no-margin pt-3 pb-2" type="text" style="width: 60px" :loading="true" v-if="loadingDataStatistics"></v-skeleton-loader>
                                    <span class="text-h6" v-if="!loadingDataStatistics">{{ displayDataStatistics ? displayDataStatistics.totalTransactionCategoryCount : '-' }}</span>
                                </div>
                            </div>
                        </v-col>
                        <v-col cols="6" sm="3">
                            <div class="d-flex align-center">
                                <div class="me-3">
                                    <v-avatar rounded color="secondary" size="42" class="elevation-1">
                                        <v-icon size="24" :icon="icons.tags"/>
                                    </v-avatar>
                                </div>

                                <div class="d-flex flex-column">
                                    <span class="text-caption">{{ $t('Transaction Tags') }}</span>
                                    <v-skeleton-loader class="skeleton-no-margin pt-3 pb-2" type="text" style="width: 60px" :loading="true" v-if="loadingDataStatistics"></v-skeleton-loader>
                                    <span class="text-h6" v-if="!loadingDataStatistics">{{ displayDataStatistics ? displayDataStatistics.totalTransactionTagCount : '-' }}</span>
                                </div>
                            </div>
                        </v-col>
                    </v-row>
                </v-card-text>
            </v-card>
        </v-col>

        <v-col cols="12" v-if="isDataExportingEnabled">
            <v-card :class="{ 'disabled': exportingData }" :title="$t('Export Data')">
                <v-card-text>
                    <span class="text-body-1">{{ $t('Export all data to csv file.') }}&nbsp;{{ $t('It may take a long time, please wait for a few minutes.') }}</span>
                </v-card-text>

                <v-card-text class="d-flex flex-wrap gap-4">
                    <v-btn :disabled="loadingDataStatistics || exportingData || !dataStatistics || !dataStatistics.totalTransactionCount || dataStatistics.totalTransactionCount === '0'" @click="exportData">
                        {{ $t('Export Data') }}
                        <v-progress-circular indeterminate size="24" class="ml-2" v-if="exportingData"></v-progress-circular>
                    </v-btn>
                </v-card-text>
            </v-card>
        </v-col>

        <v-col cols="12">
            <v-card :class="{ 'disabled': clearingData }">
                <template #title>
                    <span class="text-error">{{ $t('Danger Zone') }}</span>
                </template>

                <v-form>
                    <v-card-text class="py-0">
                    <span class="text-body-1 text-error">
                        <v-icon :icon="icons.alert"/>
                        {{ $t('You CANNOT undo this action. This will clear your accounts, categories, tags and transactions data. Please input your current password to confirm.') }}
                    </span>
                    </v-card-text>

                    <v-card-text class="pb-0">
                        <v-row class="mb-3">
                            <v-col cols="12" md="6">
                                <v-text-field
                                    ref="currentPasswordInput"
                                    autocomplete="current-password"
                                    clearable variant="underlined"
                                    color="error"
                                    :disabled="loadingDataStatistics || clearingData"
                                    :placeholder="$t('Current Password')"
                                    :type="isCurrentPasswordVisible ? 'text' : 'password'"
                                    :append-inner-icon="isCurrentPasswordVisible ? icons.eyeSlash : icons.eye"
                                    v-model="currentPasswordForClearData"
                                    @click:append-inner="isCurrentPasswordVisible = !isCurrentPasswordVisible"
                                    @keyup.enter="clearData"
                                />
                            </v-col>
                        </v-row>
                    </v-card-text>

                    <v-card-text class="d-flex flex-wrap gap-4">
                        <v-btn color="error" :disabled="loadingDataStatistics || !currentPasswordForClearData || clearingData" @click="clearData">
                            {{ $t('Clear User Data') }}
                            <v-progress-circular indeterminate size="24" class="ml-2" v-if="clearingData"></v-progress-circular>
                        </v-btn>
                    </v-card-text>
                </v-form>
            </v-card>
        </v-col>
    </v-row>

    <confirm-dialog ref="confirmDialog"/>
    <snack-bar ref="snackbar" />
</template>

<script>
import { mapStores } from 'pinia';
import { useRootStore } from '@/stores/index.js';
import { useSettingsStore } from '@/stores/setting.js';
import { useUserStore } from '@/stores/user.js';

import {appendThousandsSeparator, isEquals} from '@/lib/common.js';
import { isDataExportingEnabled } from '@/lib/server_settings.js';
import { startDownloadFile } from '@/lib/ui.js';

import {
    mdiRefresh,
    mdiListBoxOutline,
    mdiCreditCardOutline,
    mdiViewDashboardOutline,
    mdiTagOutline,
    mdiAlert,
    mdiEyeOutline,
    mdiEyeOffOutline
} from '@mdi/js';

export default {
    data() {
        return {
            loadingDataStatistics: true,
            dataStatistics: null,
            exportingData: false,
            currentPasswordForClearData: '',
            isCurrentPasswordVisible: false,
            clearingData: false,
            icons: {
                refresh: mdiRefresh,
                transactions: mdiListBoxOutline,
                accounts: mdiCreditCardOutline,
                categories: mdiViewDashboardOutline,
                tags: mdiTagOutline,
                alert: mdiAlert,
                eye: mdiEyeOutline,
                eyeSlash: mdiEyeOffOutline
            }
        }
    },
    computed: {
        ...mapStores(useRootStore, useSettingsStore, useUserStore),
        isEnableThousandsSeparator() {
            return this.settingsStore.appSettings.thousandsSeparator;
        },
        displayDataStatistics() {
            const self = this;

            if (!self.dataStatistics) {
                return null;
            }

            return {
                totalAccountCount: appendThousandsSeparator(self.dataStatistics.totalAccountCount, self.isEnableThousandsSeparator),
                totalTransactionCategoryCount: appendThousandsSeparator(self.dataStatistics.totalTransactionCategoryCount, self.isEnableThousandsSeparator),
                totalTransactionTagCount: appendThousandsSeparator(self.dataStatistics.totalTransactionTagCount, self.isEnableThousandsSeparator),
                totalTransactionCount: appendThousandsSeparator(self.dataStatistics.totalTransactionCount, self.isEnableThousandsSeparator)
            };
        },
        isDataExportingEnabled() {
            return isDataExportingEnabled();
        },
        exportFileName() {
            const nickname = this.userStore.currentUserNickname;

            if (nickname) {
                return this.$t('dataExport.exportFilename', {
                    nickname: nickname
                }) + '.csv';
            }

            return this.$t('dataExport.defaultExportFilename') + '.csv';
        }
    },
    created() {
        this.reloadUserDataStatistics(false);
    },
    methods: {
        reloadUserDataStatistics(force) {
            const self = this;

            self.loadingDataStatistics = true;

            self.userStore.getUserDataStatistics().then(dataStatistics => {
                if (force) {
                    if (isEquals(self.dataStatistics, dataStatistics)) {
                        self.$refs.snackbar.showMessage('Data is up to date');
                    } else {
                        self.$refs.snackbar.showMessage('Data has been updated');
                    }
                }

                self.dataStatistics = dataStatistics;
                self.loadingDataStatistics = false;
            }).catch(error => {
                self.loadingDataStatistics = false;

                if (!error.processed) {
                    self.$refs.snackbar.showError(error);
                }
            });
        },
        exportData() {
            const self = this;

            if (self.exportingData) {
                return;
            }

            self.exportingData = true;

            self.userStore.getExportedUserData().then(data => {
                startDownloadFile(self.exportFileName, data);
                self.exportingData = false;
            }).catch(error => {
                self.exportingData = false;

                if (!error.processed) {
                    self.$refs.snackbar.showError(error);
                }
            });
        },
        clearData() {
            const self = this;

            if (!self.currentPasswordForClearData) {
                self.$refs.snackbar.showMessage('Current password cannot be empty');
                return;
            }

            if (self.clearingData) {
                return;
            }

            self.$refs.confirmDialog.open('Are you sure you want to clear all data?', { color: 'error' }).then(() => {
                self.clearingData = true;
                self.isCurrentPasswordVisible = false;

                self.rootStore.clearUserData({
                    password: self.currentPasswordForClearData
                }).then(() => {
                    self.clearingData = false;

                    self.$refs.snackbar.showMessage('All user data has been cleared');
                    self.reloadUserDataStatistics();
                }).catch(error => {
                    self.clearingData = false;

                    if (!error.processed) {
                        self.$refs.snackbar.showError(error);
                    }
                });
            });
        }
    }
}
</script>
