import { get, writable } from "svelte/store";
import { ImageResource  } from "@models/ImageResources";
import { MEDIA_SIZES } from "@libs/Services/media_loaders";
import { browser } from "$app/environment"
/**
* @typedef {import("svelte/store").Writable<T>} Writable
 * @template T
*/

/**
 * If there is a resize on the viewport but is not bigger than this threshold, the layout is not considered changed
 * @type {number} - a value between 0 and 1
 */
const layout_change_threshold = 0.05;

let LAYOUT_PROPERTIES = {
    IS_MOBILE: false,
    MOBILE_BREAKPOINT: 768,
    TABLET_BREAKPOINT: 1024,
    VIEWPORT_WIDTH: 1920,
    VIEWPORT_HEIGHT: 1080,
    SPACING: {
        SPACING_1: 0,
        SPACING_2: 0,
        SPACING_3: 0,
        SPACING_4: 0,
        SPACING_5: 0,
        SPACING_6: 0,
        SPACING_7: 0,
        SPACING_8: 0,
        SPACING_9: 0
    }
}

if (browser) {
    LAYOUT_PROPERTIES.VIEWPORT_WIDTH = window.innerWidth;
    LAYOUT_PROPERTIES.VIEWPORT_HEIGHT = window.innerHeight;
}

/**
 * @type {Writable<typeof LAYOUT_PROPERTIES>}
 * @description the layout properties of the website
 */
export const layout_properties = writable(LAYOUT_PROPERTIES);

export const defineLayout = () => {
    const root_styles = getComputedStyle(document.documentElement);

    let new_layout_properties = {
        ...LAYOUT_PROPERTIES
    };

    new_layout_properties.SPACING = {
        VSPACING_1: root_styles.getPropertyValue("--vspacing-1"),
        VSPACING_2: root_styles.getPropertyValue("--vspacing-2"),
        VSPACING_3: root_styles.getPropertyValue("--vspacing-3"),
        VSPACING_4: root_styles.getPropertyValue("--vspacing-4"),
        VSPACING_5: root_styles.getPropertyValue("--vspacing-5"),
        VSPACING_6: root_styles.getPropertyValue("--vspacing-6"),
        VSPACING_7: root_styles.getPropertyValue("--vspacing-7"),
        VSPACING_8: root_styles.getPropertyValue("--vspacing-8"),
        VSPACING_9: root_styles.getPropertyValue("--vspacing-9")
    }

    new_layout_properties.IS_MOBILE = isMobile();
    console.log(`called, and the answer to isMobile is ${new_layout_properties.IS_MOBILE}`);

    new_layout_properties.VIEWPORT_WIDTH = window.innerWidth;
    new_layout_properties.VIEWPORT_HEIGHT = window.innerHeight;

    layout_properties.set(new_layout_properties);
    LAYOUT_PROPERTIES = new_layout_properties;
}

export const hasChangedLayout = () => {
    const current_width = window.innerWidth;
    const current_height = window.innerHeight;

    const width_change = Math.abs(current_width - LAYOUT_PROPERTIES.VIEWPORT_WIDTH) / LAYOUT_PROPERTIES.VIEWPORT_WIDTH;
    const height_change = Math.abs(current_height - LAYOUT_PROPERTIES.VIEWPORT_HEIGHT) / LAYOUT_PROPERTIES.VIEWPORT_HEIGHT;


    // const layout_changed = width_change > layout_change_threshold || height_change > layout_change_threshold;

    // return layout_changed;

    return width_change > layout_change_threshold || height_change > layout_change_threshold;
}

export function isMobile() {
    let is_mobile = /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent);
    
    if (!is_mobile && window.innerWidth < LAYOUT_PROPERTIES.MOBILE_BREAKPOINT) {
        is_mobile = true;
    }

    
    return is_mobile;
}

export const user_locale = browser ? (window.navigator.language.split("-")[0] || "en") : "en";



/*=============================================
=            Layout elements            =
=============================================*/


export const layout_images = {
    COMBINATION_MARK: new ImageResource("logos/combination_mark.webp", MEDIA_SIZES.EXTRA_LARGE.postfix),
    SELVAMAAKU_LOGOTYPE: new ImageResource("logos/selva_maaku_logo.webp", MEDIA_SIZES.MEDIUM.postfix),
    MANGO_ISOTYPE: new ImageResource("logos/mango_isotype.webp", MEDIA_SIZES.MEDIUM.postfix),
    MANGO_ISOLOGO: new ImageResource("logos/mango_isologo.webp", MEDIA_SIZES.EXTRA_LARGE.postfix),
    HERO_VIDEO_COVER: new ImageResource("covers/hero_video_cover.webp", MEDIA_SIZES.LARGE.postfix),
    KEITT_HOUSE: new ImageResource("houses/keitt_house.webp", MEDIA_SIZES.LARGE.postfix),
    KEITT_HOUSE_TWO: new ImageResource("houses/keitt_house_two.webp", MEDIA_SIZES.LARGE.postfix),
    KEITT_HOUSE_THREE: new ImageResource("houses/keitt_house_three.webp", MEDIA_SIZES.LARGE.postfix),
    KEITT_HOUSE_FOUR: new ImageResource("houses/house_keitt_four.webp", MEDIA_SIZES.LARGE.postfix),
    KEITT_HOUSE_FIVE: new ImageResource("houses/house_keitt_five.webp", MEDIA_SIZES.LARGE.postfix),
    KEITT_HOUSE_SIX: new ImageResource("houses/house_keitt_six.webp", MEDIA_SIZES.LARGE.postfix),
    TOMMY_HOUSE: new ImageResource("houses/tommy_house.webp", MEDIA_SIZES.LARGE.postfix),   
    TOMMY_HOUSE_TWO: new ImageResource("houses/tommy_house_two.webp", MEDIA_SIZES.EXTRA_LARGE.postfix),   
    TOMMY_HOUSE_THREE: new ImageResource("houses/tommy_house_three.webp", MEDIA_SIZES.EXTRA_LARGE.postfix),   
    TOMMY_HOUSE_FOUR: new ImageResource("houses/tommy_house_four.webp", MEDIA_SIZES.EXTRA_LARGE.postfix),   
    TOMMY_HOUSE_FIVE: new ImageResource("houses/house_tommy_five.webp", MEDIA_SIZES.LARGE.postfix),
    TOMMY_HOUSE_SIX: new ImageResource("houses/house_tommy_six.webp", MEDIA_SIZES.LARGE.postfix),
    COCOTEROS_HOUSE: new ImageResource("houses/house_cocoteros_one.webp", MEDIA_SIZES.EXTRA_LARGE.postfix),
    COCOTEROS_HOUSE_TWO: new ImageResource("houses/house_cocoteros_two.webp", MEDIA_SIZES.EXTRA_LARGE.postfix),
    COCOTEROS_HOUSE_THREE: new ImageResource("houses/house_cocoteros_three.webp", MEDIA_SIZES.EXTRA_LARGE.postfix),
    COCOTEROS_HOUSE_FOUR: new ImageResource("houses/house_cocoteros_four.webp", MEDIA_SIZES.EXTRA_LARGE.postfix),
    COCOTEROS_HOUSE_FIVE: new ImageResource("houses/house_cocoteros_five.webp", MEDIA_SIZES.EXTRA_LARGE.postfix),
    DEVELOPMENT_HOUSE_ONE: new ImageResource("houses/development_house_one.webp", MEDIA_SIZES.EXTRA_LARGE.postfix), 
    HOUSE_MODEL_ONE: new ImageResource("houses/house_model_one.webp", MEDIA_SIZES.MEDIUM.postfix),
    STONE_WALL: new ImageResource("misc/wixarica_wall.webp", MEDIA_SIZES.MEDIUM.postfix),
    RAINY_POOL: new ImageResource("misc/rainy_pool.webp", MEDIA_SIZES.MEDIUM.postfix),
    RAINY_POOL_TWO: new ImageResource("misc/rainy_pool_two.webp", MEDIA_SIZES.EXTRA_LARGE.postfix),
    PARK: new ImageResource("misc/selvamaaku_park.webp", MEDIA_SIZES.MEDIUM.postfix),   
    JUNGLE_ONE: new ImageResource("misc/selvamaaku_jungle.webp", MEDIA_SIZES.EXTRA_LARGE.postfix),
    SAYULITA_MAP: new ImageResource("misc/sayulita_map.webp", MEDIA_SIZES.EXTRA_LARGE.postfix),
    SAYULITA_TOWN_ONE: new ImageResource("misc/sayulita_town_one.webp", MEDIA_SIZES.EXTRA_LARGE.postfix),
    SAYULITA_TOWN_TWO: new ImageResource("misc/sayulita_town_two.webp", MEDIA_SIZES.EXTRA_LARGE.postfix),
}