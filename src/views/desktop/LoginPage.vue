<template>
    <div class="layout-wrapper">
        <router-link to="/">
            <div class="auth-logo d-flex align-start gap-x-3">
                <img alt="logo" class="login-page-logo" :src="goFireLogoPath" />
                <h1 class="font-weight-medium leading-normal text-2xl">{{ $t('global.app.title') }}</h1>
            </div>
        </router-link>
        <v-row no-gutters class="auth-wrapper">
            <v-col cols="12" md="8" class="d-none d-md-flex align-center justify-center position-relative">
                <div class="d-flex auth-img-footer" v-if="!isDarkMode">
                    <v-img src="img/desktop/background.svg"/>
                </div>
                <div class="d-flex auth-img-footer" v-if="isDarkMode">
                    <v-img src="img/desktop/background-dark.svg"/>
                </div>
                <div class="d-flex align-center justify-center w-100 pt-10">
                    <v-img max-width="600px" src="img/desktop/people1.svg"/>
                </div>
            </v-col>
            <v-col cols="12" md="4" class="auth-card d-flex flex-column">
                <div class="d-flex align-center justify-center h-100">
                    <v-card variant="flat" class="w-100 mt-0 px-4 pt-12" max-width="500">
                        <v-card-text>
                            <h5 class="text-h5 mb-3">{{ $t('Welcome to gofire') }}</h5>
                            <p class="mb-0">{{ $t('Please log in with your gofire account') }}</p>
                        </v-card-text>

                        <v-card-text class="pb-0 mb-6">
                            <v-form>
                                <v-row>
                                    <v-col cols="12">
                                        <v-text-field
                                            type="text"
                                            autocomplete="username"
                                            autofocus="autofocus"
                                            clearable
                                            :disabled="show2faInput || logining || verifying"
                                            :label="$t('Username')"
                                            :placeholder="$t('Your username or email')"
                                            v-model="username"
                                            @input="tempToken = ''"
                                            @keyup.enter="$refs.passwordInput.focus()"
                                        />
                                    </v-col>

                                    <v-col cols="12">
                                        <v-text-field
                                            autocomplete="current-password"
                                            clearable
                                            ref="passwordInput"
                                            :type="isPasswordVisible ? 'text' : 'password'"
                                            :disabled="show2faInput || logining || verifying"
                                            :label="$t('Password')"
                                            :placeholder="$t('Your password')"
                                            :append-inner-icon="isPasswordVisible ? icons.eyeSlash : icons.eye"
                                            v-model="password"
                                            @input="tempToken = ''"
                                            @click:append-inner="isPasswordVisible = !isPasswordVisible"
                                            @keyup.enter="login"
                                        />
                                    </v-col>

                                    <v-col cols="12" v-show="show2faInput">
                                        <v-text-field
                                            type="number"
                                            autocomplete="one-time-code"
                                            clearable
                                            ref="passcodeInput"
                                            :disabled="logining || verifying"
                                            :label="$t('Passcode')"
                                            :placeholder="$t('Passcode')"
                                            :append-inner-icon="icons.backupCode"
                                            v-model="passcode"
                                            @click:append-inner="twoFAVerifyType = 'backupcode'"
                                            @keyup.enter="verify"
                                            v-if="twoFAVerifyType === 'passcode'"
                                        />
                                        <v-text-field
                                            type="text"
                                            clearable
                                            :disabled="logining || verifying"
                                            :label="$t('Backup Code')"
                                            :placeholder="$t('Backup Code')"
                                            :append-inner-icon="icons.passcode"
                                            v-model="backupCode"
                                            @click:append-inner="twoFAVerifyType = 'passcode'"
                                            @keyup.enter="verify"
                                            v-if="twoFAVerifyType === 'backupcode'"
                                        />
                                    </v-col>

                                    <v-col cols="12" class="py-0 mt-1 mb-4">
                                        <div class="d-flex align-center justify-space-between flex-wrap">
                                            <a href="javascript:void(0);" @click="showMobileQrCode = true">
                                                <span class="nav-item-title">{{ $t('Use on Mobile Device') }}</span>
                                            </a>
                                            <v-spacer/>
                                            <router-link class="text-primary" to="/forgetpassword"
                                                         :class="{'disabled': !isUserForgetPasswordEnabled}">
                                                {{ $t('Forget Password?') }}
                                            </router-link>
                                        </div>
                                    </v-col>

                                    <v-col cols="12">
                                        <v-btn block :disabled="inputIsEmpty || logining || verifying"
                                               @click="login" v-if="!show2faInput">
                                            {{ $t('Log In') }}
                                            <v-progress-circular indeterminate size="24" class="ml-2" v-if="logining"></v-progress-circular>
                                        </v-btn>
                                        <v-btn block :disabled="twoFAInputIsEmpty || logining || verifying"
                                               @click="verify" v-else-if="show2faInput">
                                            {{ $t('Continue') }}
                                            <v-progress-circular indeterminate size="24" class="ml-2" v-if="verifying"></v-progress-circular>
                                        </v-btn>
                                    </v-col>

                                    <v-col cols="12" class="text-center text-base">
                                        <span class="me-1">{{ $t('Don\'t have an account?') }}</span>
                                        <router-link class="text-primary" to="/signup"
                                                     :class="{'disabled': !isUserRegistrationEnabled}">
                                            {{ $t('Create an account') }}
                                        </router-link>
                                    </v-col>
                                </v-row>
                            </v-form>
                        </v-card-text>
                    </v-card>
                </div>
                <v-spacer/>
                <div class="d-flex align-center justify-center">
                    <v-card variant="flat" class="w-100 px-4 pb-4" max-width="500">
                        <v-card-text class="pt-0">
                            <v-row>
                                <v-col cols="12" class="text-center">
                                    <v-menu location="bottom">
                                        <template #activator="{ props }">
                                            <v-btn variant="text"
                                                   :disabled="logining || verifying"
                                                   v-bind="props">{{ currentLanguageName }}</v-btn>
                                        </template>
                                        <v-list>
                                            <v-list-item v-for="(lang, locale) in allLanguages" :key="locale">
                                                <v-list-item-title
                                                    class="cursor-pointer"
                                                    @click="changeLanguage(locale)">
                                                    {{ lang.displayName }}
                                                </v-list-item-title>
                                            </v-list-item>
                                        </v-list>
                                    </v-menu>
                                </v-col>

                                <v-col cols="12" class="d-flex align-center pt-0">
                                    <v-divider />
                                </v-col>

                                <v-col cols="12" class="text-center text-sm">
                                    <span>Powered by </span>
                                    <a href="https://github.com/f97/gofire" target="_blank">gofire</a>&nbsp;<span>{{ version }}</span>
                                </v-col>
                            </v-row>
                        </v-card-text>
                    </v-card>
                </div>
            </v-col>
        </v-row>

        <switch-to-mobile-dialog v-model:show="showMobileQrCode" />
        <snack-bar ref="snackbar" />
    </div>
</template>

<script>
import { useTheme } from 'vuetify';

import { mapStores } from 'pinia';
import { useRootStore } from '@/stores/index.js';
import { useSettingsStore } from '@/stores/setting.js';
import { useExchangeRatesStore } from '@/stores/exchangeRates.js';

import assetConstants from '@/consts/asset.js';
import { isUserRegistrationEnabled, isUserForgetPasswordEnabled } from '@/lib/server_settings.js';

import {
    mdiEyeOutline,
    mdiEyeOffOutline,
    mdiOnepassword,
    mdiHelpCircleOutline
} from '@mdi/js';

export default {
    data() {
        return {
            username: '',
            password: '',
            passcode: '',
            backupCode: '',
            tempToken: '',
            isPasswordVisible: false,
            logining: false,
            verifying: false,
            show2faInput: false,
            twoFAVerifyType: 'passcode',
            showMobileQrCode: false,
            icons: {
                eye: mdiEyeOutline,
                eyeSlash: mdiEyeOffOutline,
                passcode: mdiOnepassword,
                backupCode: mdiHelpCircleOutline
            }
        };
    },
    computed: {
        ...mapStores(useRootStore, useSettingsStore, useExchangeRatesStore),
        goFireLogoPath() {
            return assetConstants.goFireLogoPath;
        },
        version() {
            return 'v' + this.$version;
        },
        allLanguages() {
            return this.$locale.getAllLanguageInfos();
        },
        isUserRegistrationEnabled() {
            return isUserRegistrationEnabled();
        },
        isUserForgetPasswordEnabled() {
            return isUserForgetPasswordEnabled();
        },
        inputIsEmpty() {
            return !this.username || !this.password;
        },
        twoFAInputIsEmpty() {
            if (this.twoFAVerifyType === 'backupcode') {
                return !this.backupCode;
            } else {
                return !this.passcode;
            }
        },
        isDarkMode() {
            return this.globalTheme.global.name.value === 'dark';
        },
        currentLanguageName() {
            return this.$locale.getCurrentLanguageDisplayName();
        }
    },
    setup() {
        const theme = useTheme();

        return {
            globalTheme: theme
        };
    },
    methods: {
        login() {
            const self = this;

            if (!self.username) {
                self.$refs.snackbar.showMessage('Username cannot be empty');
                return;
            }

            if (!self.password) {
                self.$refs.snackbar.showMessage('Password cannot be empty');
                return;
            }

            if (self.tempToken) {
                self.show2faInput = true;
                return;
            }

            if (self.logining) {
                return;
            }

            self.isPasswordVisible = false;
            self.logining = true;

            self.rootStore.authorize({
                loginName: self.username,
                password: self.password
            }).then(authResponse => {
                self.logining = false;

                if (authResponse.need2FA) {
                    self.tempToken = authResponse.token;
                    self.show2faInput = true;

                    self.$nextTick(() => {
                        if (self.$refs.passcodeInput) {
                            self.$refs.passcodeInput.focus();
                            self.$refs.passcodeInput.select();
                        }
                    });

                    return;
                }

                if (authResponse.user) {
                    const localeDefaultSettings = self.$locale.setLanguage(authResponse.user.language);
                    self.settingsStore.updateLocalizedDefaultSettings(localeDefaultSettings);
                }

                if (self.settingsStore.appSettings.autoUpdateExchangeRatesData) {
                    self.exchangeRatesStore.getLatestExchangeRates({ silent: true, force: false });
                }

                self.$router.replace('/');
            }).catch(error => {
                self.logining = false;

                if (error.error && error.error.errorCode === 201020 && error.error.context && error.error.context.email) {
                    self.$router.push('/verify_email?email=' + encodeURIComponent(error.error.context.email));
                    return;
                }

                if (!error.processed) {
                    self.$refs.snackbar.showError(error);
                }
            });
        },
        verify() {
            const self = this;

            if (self.twoFAInputIsEmpty || self.verifying) {
                return;
            }

            if (self.twoFAVerifyType === 'passcode' && !self.passcode) {
                self.$refs.snackbar.showMessage('Passcode cannot be empty');
                return;
            } else if (self.twoFAVerifyType === 'backupcode' && !self.backupCode) {
                self.$refs.snackbar.showMessage('Backup code cannot be empty');
                return;
            }

            self.verifying = true;

            self.rootStore.authorize2FA({
                token: self.tempToken,
                passcode: self.twoFAVerifyType === 'passcode' ? self.passcode : null,
                recoveryCode: self.twoFAVerifyType === 'backupcode' ? self.backupCode : null
            }).then(authResponse => {
                self.verifying = false;

                if (authResponse.user) {
                    const localeDefaultSettings = self.$locale.setLanguage(authResponse.user.language);
                    self.settingsStore.updateLocalizedDefaultSettings(localeDefaultSettings);
                }

                if (self.settingsStore.appSettings.autoUpdateExchangeRatesData) {
                    self.exchangeRatesStore.getLatestExchangeRates({ silent: true, force: false });
                }

                self.$router.replace('/');
            }).catch(error => {
                self.verifying = false;

                if (!error.processed) {
                    self.$refs.snackbar.showError(error);
                }
            });
        },
        changeLanguage(locale) {
            const localeDefaultSettings = this.$locale.setLanguage(locale);
            this.settingsStore.updateLocalizedDefaultSettings(localeDefaultSettings);
        }
    }
}
</script>
