import { writable } from "svelte/store";

export const supported_languages = {
    ENGLISH: 'en',
    SPANISH: 'es',
}

/**
 * The user's preferred language that is supported by the site. YOU SHOULD NOT SET THIS DIRECTLY, instead use the setSiteLanguage function.
 * @type {import('svelte/store').Writable<string>}
 */
export const site_language = writable('en');

/**
 * Sets the site language to the provided language if it is supported.
 * @param {string} locale
 * @returns {void}
 */
export const setSiteLanguage = locale => {
    for (let lang of Object.values(supported_languages)) {
        if (locale.includes(lang)) {
            site_language.set(lang);
            break;
        }
    }
}