<script>
    import InteractiveImage from "@components/Images/InteractiveImage.svelte";
import viewport from "@components/viewport_actions/useViewportActions";
    import { layout_images } from "@stores/layout";
    import { setRenderMedia } from "@stores/media_display";
    import { elasticIn, elasticOut } from "svelte/easing";
    import { fly } from "svelte/transition";

    
    /*=============================================
    =            Properties            =
    =============================================*/

        /**
         * Whether the parent component is visible or not
         * @type {import("svelte/store").Writable<boolean>}
         */
        export let parent_visible;

        /**
         * Whether the gallery is visible or not
         * @type {boolean}
         */
        let gallery_visible = false;

        /**
         * Duration of the entry animation in milliseconds
         * @type {number}
         */
        export let entry_animation_duration = 1500;

        const fly_transition_params = {
            delay: entry_animation_duration * 0.1,
            duration: entry_animation_duration,
            y: -600,
            x: 0,
            opacity: 0,
            easing: elasticOut
        }
    
    /*=====  End of Properties  ======*/

    
    /*=============================================
    =            Methods            =
    =============================================*/
    
        const setGalleryVisibility = is_visible => {
            gallery_visible = is_visible;
        }
    
    /*=====  End of Methods  ======*/
    
    
        
</script>

<div
    id="amenities-gallery-wrapper"
    style:--entry-animation-duration="{entry_animation_duration}ms"
    on:viewportEnter={() => setGalleryVisibility(true)}
    use:viewport={{ height_offset: 0.4 }}
>
    {#if gallery_visible && $parent_visible}
        <!-- TODO: Add alt text to these images -->
        <div id="ag-park" 
            class="amenities-gallery-image"
            transition:fly={{
                ...fly_transition_params,
                y: fly_transition_params.y * -1,
            }}
        >
            <InteractiveImage 
                on:image-clicked={() => setRenderMedia(layout_images.PARK.getUrl(1))}
                image_resource={layout_images.PARK}
                desktop_viewport_percentage={0.3}
                mobile_viewport_percentage={0.3}
                overlay_color="var(--theme-color)"
                max_overlay_opacity={0.5}
            />
        </div>
        <div id="ag-rainy-pool" 
            class="amenities-gallery-image"
            transition:fly={{
                ...fly_transition_params,
                y: fly_transition_params.y * 1.5,
            }}
        >
            <InteractiveImage 
                on:image-clicked={() => setRenderMedia(layout_images.RAINY_POOL.getUrl(1))}
                image_resource={layout_images.RAINY_POOL}
                desktop_viewport_percentage={0.3}
                mobile_viewport_percentage={0.3}
                overlay_color="var(--theme-color)"
                max_overlay_opacity={0.5}
            />
        </div>
        <div id="ag-stone-wall" 
            class="amenities-gallery-image"
            transition:fly={{ 
                ...fly_transition_params, 
                y: fly_transition_params.y,
                delay: fly_transition_params.delay * 0.4
            }}
        >
            <InteractiveImage 
                on:image-clicked={() => setRenderMedia(layout_images.STONE_WALL.getUrl(1))}
                image_resource={layout_images.STONE_WALL}
                desktop_viewport_percentage={0.3}
                mobile_viewport_percentage={0.3}
                overlay_color="var(--theme-color)"
                max_overlay_opacity={0.5}
            />
        </div>
    {/if}
</div>

<style>
    #amenities-gallery-wrapper {
        position: relative;
        width: min(811px, 66.69407cqw);
        height: min(369px, 30.34539cqw);
    }

    .amenities-gallery-image {
        position: absolute;
        object-fit: cover;
        object-position: center;
    }

    
    /*=============================================
    =            Elements positions            =
    =============================================*/
    
        #ag-park {
            left: 0;
            top: 16.26%;
            width: 24.90752%;
            height: 83.46883%;
            z-index: var(--z-index-t-2);
        }

        #ag-rainy-pool {
            top: 0;
            left: 20.22%;
            width: 59.92601%;
            height: 86.17886%;
            z-index: var(--z-index-t-1);
        }

        #ag-stone-wall {
            right: 0;
            top: 16.26%;
            width: 22.81134%;
            height: 53.38753%;
            z-index: var(--z-index-t-2);
        }
    
    /*=====  End of Elements positions  ======*/
    
    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @container (width <= 768px) {
            #amenities-gallery-wrapper {
                width: 100cqw;
                height: 45.49938cqw;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
    
</style>