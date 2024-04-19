import { writable } from "svelte/store";

/**
 * The url of media that will be displayed by the media viewer. If its an empty string then the media viewer will not be rendered.
 * @type {import("svelte/store").Writable<string>}
 * @default ''
 */
export const render_media = writable("");


/**
 * Sets the render media url.
 * @param {string} url
 */
export const setRenderMedia = (url) => render_media.set(url);

/**
 * Resets the render media url.
 */
export const resetRenderMedia = () => render_media.set("");