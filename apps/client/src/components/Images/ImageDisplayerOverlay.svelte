<script>
    import { render_media, resetRenderMedia } from "@stores/media_display";
    import { onDestroy, onMount, tick } from "svelte";
    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        /** 
         * Whether the media has loaded.
         * @type {boolean}
         */
        let is_loaded = false;

        /**
         * Media display
         * @type {HTMLElement}
        */
        let media_display;

        /**
         * @typedef {Object} Point
         * @property {number} x
         * @property {number} y
        */

        /**
         * The coordinate of the close button.
         * @type {Point}
         * @default null
         */
        let close_button_position = null;

        /**
         * The close btn size
         * @type {number}
         */
        const close_btn_size = 50;


        let render_media_change_unsubscriber;
    
    /*=====  End of Properties  ======*/

    onMount(() => {
        if ($render_media !== "") {
            loadNewMedia($render_media);
        }

        render_media_change_unsubscriber = render_media.subscribe(loadNewMedia);
    });

    onDestroy(() => {
        if (render_media_change_unsubscriber != null) {
            render_media_change_unsubscriber();
        }
    })

    
    /*=============================================
    =            Methods            =
    =============================================*/
    
        /**
         * Handles the load of the current render media.
         * @param {string} media_url
        */
       const loadNewMedia = media_url => {
            if (media_url === "") return;

            const media_element = document.createElement('img');

            media_element.onload = async () => {
                is_loaded = true;
                // await tick();
                const display = media_display.querySelector('img');
                display.src = media_url;
                await tick();
                display.onload = calcCloseBtnPosition;
                // calcCloseBtnPosition();
            }

            media_element.src = media_url;
       }

       /**
        * Calculates where the close button should be positioned so it's on the top right corner of the media.
        */
       const calcCloseBtnPosition = () => {
            if (media_display == null) return;

            const media_display_styles = window.getComputedStyle(media_display);
            const media_display_padding_top = parseInt(media_display_styles.getPropertyValue('padding-top'));
            const media_display_padding_left = parseInt(media_display_styles.getPropertyValue('padding-left'));
            console.log(media_display_styles);

            const media = media_display.querySelector('img');
            const media_rect = media.getBoundingClientRect();
            console.log(media_rect);

            close_button_position = {
                x: (media_rect.x + media_rect.width - (close_btn_size * 0.5)) - media_display_padding_left,
                y: (media_rect.y - (close_btn_size * 0.5)) - media_display_padding_top
            }
       }

       const handleCloseBtnClick = () => {
           resetRenderMedia();
           is_loaded = false;
       }
    
    /*=====  End of Methods  ======*/
    
    
    
    
</script>

{#if $render_media !== ""}
    <article id="selected-media-render"
        bind:this={media_display}
    >
        <div id="smr-content-wrapper" class:dtwo={false}>
            {#if close_button_position != null}
                <button class="button-2" 
                    on:click={handleCloseBtnClick}
                    style:--close-btn-size={close_btn_size + "px"}
                    style:top={close_button_position.y + "px"}
                    style:left={close_button_position.x + "px"}
                >
                    <svg viewBox="0 0 50 50">
                        <path d="M3 3L47 47M3 47L47 3"></path>
                    </svg>
                </button>
            {/if}
            {#if is_loaded}
                <img 
                    src="{$render_media}" 
                    alt="" 
                    id="selected-media"
                >
            {:else}
                <div id="smk-smr-loader"></div>
            {/if}
        </div>
    </article>
{/if}

<style>
    #selected-media-render {
        position: fixed;
        background-color: hsla(158, 17%, 25%, 0.589);
        top: 0;
        left: 0;
        width: 100vw;
        height: 100dvh;
        z-index: var(--z-index-t-4);
        padding: calc(var(--spacing-4) + var(--navbar-height)) var(--spacing-4) var(--spacing-4) var(--spacing-4);
    }

    #smr-content-wrapper {
        position: relative;
        width: 100%;
        height: 100%;
        display: grid;
        place-items: center;
    }

    #smr-content-wrapper button.button-2 {
        --close-btn-size: 40px;

        position: absolute;
        height: var(--close-btn-size);
        width: var(--close-btn-size);
        border-radius: 50%;
        padding: calc(var(--close-btn-size) * 0.25);
        background-color: var(--theme-color);
        transition: background-color 0.8s ease;

        &:hover {
            background-color: var(--theme-color-light);
        }

        & svg  {
            width: 100%;
            height: 100%;
            fill: none;
        }

        & svg path {
            stroke: var(--light-orange-light-active);
            stroke-width: calc(var(--close-btn-size) * 0.1);
            stroke-linecap: round;
        }
    }

    @supports (color: rgb(from white r g b)) {
        #selected-media-render {
            background-color: hsl(from var(--body-bg-color) h s l / 0.9);
        }

        #smr-content-wrapper button.button-2 {
            background-color: hsl(from var(--theme-color) h 100% l / 0.6);
        }

        #smr-content-wrapper button.button-2:hover {
            background-color: hsl(from var(--theme-color) h 100% calc(l * 1.5) / 0.8);
        }
    }

    #selected-media {
        max-width: 80%;
        max-height: 80%;
    }

    #smk-smr-loader {
        --loader-size: 80px;

        width: var(--loader-size);
        height: var(--loader-size);
        color: var(--light-orange-light-active);
        will-change: width,height;
        background:
            conic-gradient(from  -45deg at top    calc(var(--loader-size) * 0.5) left 50% ,#0000 ,currentColor 1deg 90deg,#0000 91deg),
            conic-gradient(from   45deg at right  calc(var(--loader-size) * 0.5) top  50% ,#0000 ,currentColor 1deg 90deg,#0000 91deg),
            conic-gradient(from  135deg at bottom calc(var(--loader-size) * 0.5) left 50% ,#0000 ,currentColor 1deg 90deg,#0000 91deg),
            conic-gradient(from -135deg at left   calc(var(--loader-size) * 0.5) top  50% ,#0000 ,currentColor 1deg 90deg,#0000 91deg);
        animation: loader-animation 1.8s infinite cubic-bezier(0.3,1,0,1);
    }

    @keyframes loader-animation {
        50%  {width: calc(var(--loader-size) * 1.5);height: calc(var(--loader-size) * 1.5);transform: rotate(180deg)}
        100% {transform: rotate(360deg)}
    }
</style>