<script>
    import { layout_images, layout_properties } from "@stores/layout";
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import HeroVideo from "../sub-components/hero_video.svelte";

    
    /*=============================================
    =            Properties            =
    =============================================*/
    
    
        /*----------  Animation  ----------*/

            /**
             * The particles wrapper
             * @type {HTMLElement}
             */
            let particles_wrapper;
    
            /**
             * determines how much the particles will move in relation to the mouse
             * @type {number}
             */
            const mouse_tracking_factor = 10;            


            /**
             * Whether the mouse position variable is locked
             * @type {boolean}
             */
            let mouse_position_locked = false;

            /**
             * The duration of the mouse position lock
             * @type {number}
             */
            const mouse_position_lock_duration = 1200;

            /**
             * Catch for the mouse position when is locked
             */
            let mouse_cache_x = 0;
            let mouse_cache_y = 0;

            let mouse_x = 0; // in %
            let mouse_y = 0;
            
    
    /*=====  End of Properties  ======*/

    
    /*=============================================
    =            Methods            =
    =============================================*/

        /**
         * Updates the mouse position 
         * @param {MouseEvent} event 
         */
        const updateMousePosition = (event) => {
            if (particles_wrapper == null || $layout_properties.IS_MOBILE) return;

            const wrapper_rect = particles_wrapper.getBoundingClientRect();

            let x_in_px = event.clientX - wrapper_rect.left;
            let y_in_px = event.clientY - wrapper_rect.top;

            mouse_cache_x =(x_in_px / wrapper_rect.width) * mouse_tracking_factor;
            mouse_cache_y = (y_in_px / wrapper_rect.height) * mouse_tracking_factor;

            if (!mouse_position_locked) {
                mouse_x = mouse_cache_x;
                mouse_y = mouse_cache_y;

                setTimeout(() => {
                    mouse_x = mouse_cache_x;
                    mouse_y = mouse_cache_y;
                    mouse_position_locked = false;
                }, mouse_position_lock_duration);

                mouse_position_locked = true;
            }

        }
    
    /*=====  End of Methods  ======*/
    
    
    
    
</script>

<section id="smk-hero-section" class:adebug={false} bind:this={particles_wrapper} on:mousemove={updateMousePosition}>
    <div class="wixarica-underlay">
        <div 
            class="smk-hp-particle-wrapper"
            style:transform="translate({-mouse_x * 0.2}%, {-mouse_y * 0.2}%)"
            style:transition="transform {mouse_position_lock_duration}ms ease-in-out"
        >
            <WixaricaIcon opacity={0.3}/>
        </div>
    </div>
    <header id="smk-hp-header" on:mousemove={updateMousePosition}>
        <div id="combination-mark-wrapper">
            <img src="{layout_images.COMBINATION_MARK.getUrl($layout_properties.IS_MOBILE ? 0.5 : 0.13)}" alt="selvamaaku combination mark">
        </div>
        <div id="smk-hero-headline-wrapper">
            <div 
                class="wixarica-decoration smk-hp-particle-wrapper"
                style:transform="translate({-mouse_x}%, {-mouse_y}%)"
                style:transition="transform {mouse_position_lock_duration}ms ease-in-out"
                class:hide-on-mobile={$layout_properties.IS_MOBILE}
            >
                <WixaricaIcon opacity={0.4}/>
            </div>
            <h1 id="smk-hero-headline">
                Living where Jungle Bliss<br/>Meets Coastal Luxury
            </h1>
            <div 
                class="wixarica-decoration smk-hp-particle-wrapper"
                style:transform="translate({-mouse_x}%, {-mouse_y}%)"
                style:transition="transform {mouse_position_lock_duration}ms ease-in-out"
                class:hide-on-mobile={$layout_properties.IS_MOBILE}
            >
                <WixaricaIcon opacity={0.4}/>
            </div>
        </div>
        <div id="smk-hero-subheadline-wrapper">
            <p id="smk-hero-subheadline">
                Live in the most exclusive house development in the jungle outside Sayulita, Mexico.
            </p>
        </div>
    </header>
    <div id="hero-video-wrapper" class:adebug={false} on:mousemove={updateMousePosition}>
        <HeroVideo
            video_width={!$layout_properties.IS_MOBILE ? "71.11cqw" : "91.7948cqw"}
            video_height={!$layout_properties.IS_MOBILE ? "40cqw" : "72.2223cqw"}
        />
    </div>
</section>

<style>
    header#smk-hp-header {
        position: relative;
        display: flex;
        flex-direction: column;
        padding: var(--spacing-4) 0 0 0;
        row-gap: var(--spacing-3);
        align-items: center;
    }

    #smk-hero-section .wixarica-underlay {
        position: absolute;
        top: 0%;
        left: 50%;
        width: 100%;
        max-width: var(--design-content-width);
        z-index: var(--z-index-b-3);
        transform: translateX(-50%);
        overflow: visible;
    }

    #smk-hero-section .wixarica-underlay .smk-hp-particle-wrapper {
        mask: linear-gradient(0deg, white 0%, white 50%, #ffffff22 100%);
    }

    
    /*=============================================
    =            Combination mark            =
    =============================================*/
    
        #combination-mark-wrapper {
            width: 100%;
            display: grid;
            place-items: center;
        }

        #combination-mark-wrapper img {
            width: 242px;
        }
    
    /*=====  End of Combination mark  ======*/

    
    /*=============================================
    =            Headline            =
    =============================================*/
    
        #smk-hero-headline-wrapper {
            display: flex;
            width: max-content;
            gap: var(--spacing-4);
            align-items: center;
            justify-content: center;
        }

        .wixarica-decoration {
            width: 120px;
        }

        h1#smk-hero-headline {
            font-size: var(--font-size-h2);
            text-align: center;
        }

        #smk-hero-subheadline-wrapper {
            width: max-content;
        } 

        p#smk-hero-subheadline {
            max-width: 527px;
            font-family: var(--font-read);
            font-size: var(--font-size-2);
            text-align: center;
        }
    
    /*=====  End of Headline  ======*/
    
    
    /*=============================================
    =            Video Section            =
    =============================================*/
    
        #hero-video-wrapper {
            display: flex;
            container-type: inline-size;
            width: min(100%, var(--design-width));
            justify-content: center;
            padding: var(--spacing-4) 0;
            margin: 0 auto;
        }
    
    /*=====  End of Video Section  ======*/
    
    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 768px) {
            #smk-hero-section .wixarica-underlay {
                width: 222vw;
                max-width: none;
            }

            #combination-mark-wrapper img {
                width: 51.282dvw;
            }

            #smk-hero-headline-wrapper {
                width: 100%;
            }

            h1#smk-hero-headline {
                font-size: calc(var(--font-size-h1) * 0.8); 
            }

            #smk-hero-subheadline-wrapper {
                width: auto;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
    
    
</style>