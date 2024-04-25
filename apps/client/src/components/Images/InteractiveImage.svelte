<script>
    import { ImageResource } from "@models/ImageResources";
    import { layout_properties } from "@stores/layout";
    import { createEventDispatcher, onMount } from "svelte";

    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        /**
         * The image resource object.
         * @type {ImageResource}
         */
        export let image_resource;

        /**
         * The image alt text.
         * @type {string}
         */
        export let alt_text = 'Interactive image';

        /**
         * The id of the image.
         * @type {string}
         */
        export let image_id = `interactive-image-${crypto.randomUUID()}`;

        /**
         * Overlay color
         * @type {string}
         * @default "rgba(0, 0, 0, 1)"
         */
        export let overlay_color = "rgba(0, 0, 0, 1)";

        /**
         * Max overlay opacity
         * @type {number}
         * @default 0.8
        */
        export let max_overlay_opacity = 0.8;
        
        /*----------  Animations  ----------*/
        
            /**
             * turns to true for the duration of the animation whenever the overlay is clicked
             * @type {boolean}
            */
            let was_clicked = false;

            /**
             * The duration of the animation
             * @type {number}
            */
            let animation_duration = 200;

            const dispatch = createEventDispatcher();

        /*----------  Viewport percentages  ----------*/

            /**
             * The image ideal viewport percentage in desktop.
             * @type {number}
             * @default 1
            */
            export let desktop_viewport_percentage = 1;

            /**
             * The image ideal viewport percentage in tablet.
             * @type {number}
             * @default 1
            */
            export let tablet_viewport_percentage = 1;

            /**
             * The image ideal viewport percentage in mobile.
             * @type {number}
             * @default 1
            */
            export let mobile_viewport_percentage = 1;

            /**
             * The image percentage to be used on the current viewport size 
             * @type {number}
             */
            let current_viewport_percentage = desktop_viewport_percentage;
    
    /*=====  End of Properties  ======*/

    onMount(() => {
        if ($layout_properties.IS_MOBILE) {
            current_viewport_percentage = mobile_viewport_percentage;
        }
    })

    
    /*=============================================
    =            Methods            =
    =============================================*/
    
        const toggleAnimationState = () => {
            was_clicked = true;
            setTimeout(() => {
                was_clicked = false;
                handleOverlayClick();
            }, animation_duration * 1.5);
        }

        const handleOverlayClick = () => {
            dispatch("image-clicked");
        }
    
    /*=====  End of Methods  ======*/
    
    
    
</script>

<div class="libery-interactive-image-wrapper">
    <div class="liiw-image-overlay" 
        style:background-color={overlay_color}
        style:--max-overlay-opacity={max_overlay_opacity}
        on:click={toggleAnimationState}
    ></div>
    <img class="liiw-image"
        class:animation-toggle={was_clicked}
        style:animation-duration="{animation_duration}ms"
        id="{image_id}"
        src="{image_resource.getUrl(current_viewport_percentage)}" 
        alt="{alt_text}" 
    >
</div>

<style>
    .libery-interactive-image-wrapper {
        position: relative;
        width: 100%;
        height: 100%;
    }

    .liiw-image-overlay {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        opacity: 0;
        transition: opacity 0.3s ease;
    }

    .animation-toggle {
        animation-name: image-interaction;
        animation-timing-function: cubic-bezier(.43,.59,.94,.64);
        animation-fill-mode: both;
        animation-direction: alternate;
        animation-iteration-count: 2;
    }

    @media(pointer:fine) {
        .liiw-image-overlay:hover {
            opacity: var(--max-overlay-opacity);
        }
    }

    @keyframes image-interaction {
        to {
            transform: scale(1.04);
        }
    }
</style>