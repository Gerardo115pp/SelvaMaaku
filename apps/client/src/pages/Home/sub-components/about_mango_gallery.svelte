<script>
    import { layout_images } from "@stores/layout";
    import viewport from "@components/viewport_actions/useViewportActions";
    import { fade } from "svelte/transition";
    import { browser } from "$app/environment";

    
    /*=============================================
    =            Properties            =
    =============================================*/

        /**
         * The Gallery component
         * @type {HTMLElement}
         */
        let gallery_component = null;
    
        
        /*----------  Animation  ----------*/
        
            /**
             * True if the gallery is visible
             * @type {boolean}
             */
            let gallery_visible = false;

            /**
             * The gallery scroll progress
             * @type {number}
             */
            let gallery_scroll_progress = 0;

            /**
             * Animated components css class
             * @type {string}
             */
            const animated_component_class = "libery-animation-component"

            /**
             * @typedef {Object} Point
             * @property {number} x
             * @property {number} y
            */

            /**
             * @typedef {Object} AnimationFrame
             * @property {Point} position
             * @property {number} scale
             * @property {number} opacity
             * @property {number} progress_trigger
            */

            /**
             * The animation frames for first image
             * @type {AnimationFrame[]}
             */
            const keitt_house_animation_frames = [
                {
                    position: {x: 0.5, y: 0.05},
                    scale: 1,
                    opacity: 0.2,
                    progress_trigger: 0.1
                },
                {
                    position: {x: 0.1, y: 0},
                    scale: 1,
                    opacity: 0.5,
                    progress_trigger: 0.25
                },
                {
                    position: {x: 0.0, y: 0},
                    scale: 1,
                    opacity: 1,
                    progress_trigger: 0.4
                }
            ]

            /**
             * The animation frames for second image
             * @type {AnimationFrame[]}
             */
            const tommy_house_animation_frames = [
                {
                    position: {x: -0.25, y: 0.25},
                    scale: 1,
                    opacity: 0.2,
                    progress_trigger: 0.05
                },
                {
                    position: {x: -0.1, y: 0.1},
                    scale: 1,
                    opacity: 0.5,
                    progress_trigger: 0.22
                },
                {
                    position: {x: 0.0, y: 0},
                    scale: 1,
                    opacity: 1,
                    progress_trigger: 0.4
                }
            ];

            /**
             * The animation frames for third image
             * @type {AnimationFrame[]}
             */
            const development_house_one_animation_frames = [
                {
                    position: {x: -0.3, y: -0.4},
                    scale: 0.7,
                    opacity: 0,
                    progress_trigger: 0.05
                },
                {
                    position: {x: -0.2, y: -0.39},
                    scale: 0.9,
                    opacity: 0.5,
                    progress_trigger: 0.2
                },
                {
                    position: {x: -0.1, y: -0.1},
                    scale: 1,
                    opacity: 1,
                    progress_trigger: 0.3
                },
                {
                    position: {x: 0, y: 0},
                    scale: 1,
                    opacity: 1,
                    progress_trigger: 0.4
                }
            ];

            /**
             * @typedef {Object} AnimatedComponent
             * @property {HTMLElement} component
             * @property {AnimationFrame[]} animation_frames
             */

            /**
             * The animated components
             * @type {AnimatedComponent[]}
             */
            let animated_components = [];
            
        

    
    /*=====  End of Properties  ======*/

    
    /*=============================================
    =            Methods            =
    =============================================*/

        /**
         * Applies an animation frame to a component
         * @param {HTMLElement} component
         * @param {AnimationFrame} animation_frame
        */
        const applyAnimationFrame = (component, animation_frame) => {
            if (!(component instanceof HTMLElement)) return;

            component.style.transform = `translate(${animation_frame.position.x * 100}%, ${animation_frame.position.y * 100}%) scale(${animation_frame.scale})`;
            component.style.opacity = animation_frame.opacity;
        }

        /**
         * Animates the components passed as arguments based on a scroll progress also passed as argument
         * @param {AnimatedComponent[]} animated_components
         * @param {number} scroll_progress
         * @returns {void}
         */
        const animateComponents = (animated_components, scroll_progress) => {
            for (let ac of animated_components) {
                let component_last_frame = parseFloat(ac.component.getAttribute('data-last-frame'));
                let max_triggered_frame = {progress_trigger: -Infinity};

                for (let frame of ac.animation_frames) {
                    if (scroll_progress >= frame.progress_trigger && frame.progress_trigger > max_triggered_frame.progress_trigger) {
                        max_triggered_frame = frame;
                    }
                }
                
                if (max_triggered_frame.position != null) {
                    applyAnimationFrame(ac.component, max_triggered_frame);
                    ac.component.setAttribute('data-last-frame', max_triggered_frame.progress_trigger);
                }
            }
        }

        /**
         * Handles the scroll event while the gallery is visible
         * @returns {void}
         */
        const handleGalleryScroll = () => {
            if (!browser) return;

            const gallery_dom_rect = gallery_component.getBoundingClientRect();

            // the distance from the origin of the document to the first pixel of the gallery
            const gallery_top = gallery_component.getBoundingClientRect().top + window.scrollY;

            // the scrolling point where the scroll progress becomes 100%
            const max_scroll_distance = gallery_top + gallery_dom_rect.height;

            gallery_scroll_progress = Math.min(1, Math.max(0, Math.min(max_scroll_distance, window.scrollY) / max_scroll_distance));
            // console.debug("Gallery scroll progress: ", gallery_scroll_progress);

            animateComponents(animated_components, gallery_scroll_progress);
        }

        /**
         * Registers an animated component
         * @param {HTMLElement} component
         * @param {AnimationFrame[]} animation_frames
         * @returns {void}
         */
        const registerAnimatedComponent = (component, animation_frames) => {
            if (!(component instanceof HTMLElement)) {
                let err = new Error("The component passed is not an HTMLElement");
                console.error(err);
                throw err;
            };

            animated_components.push({component, animation_frames});

            component.classList.add(animated_component_class);

            return {
                destroy() {
                    console.debug("Destroying animated component");
                }
            }
        }

        /**
         * Cleans up all the gallery animation data
         * @returns {void}
         */
        const resetGalleryState = () => {
            gallery_scroll_progress = 0;
            animated_components = [];
        }
        
    
        /**
         * Sets the gallery visibility
         * @param {boolean} visibility
         * @returns {void}
         */
        const setGalleryVisibility = (visibility) => {
            gallery_visible = visibility;

            if (visibility && gallery_component != null) {
                document.addEventListener('scroll', handleGalleryScroll, {passive: true});
            } else {
                document.removeEventListener('scroll', handleGalleryScroll, {passive:true});
                resetGalleryState();
            }
        }
        
    
    /*=====  End of Methods  ======*/
    
    
    
    
</script>

<div
    id="smk-mango-gallery"
    bind:this={gallery_component}
    on:viewportEnter={() => setGalleryVisibility(true)}
    on:viewportLeave={() => setGalleryVisibility(false)}
    use:viewport
>
    {#if gallery_visible}
        <img
            src={layout_images.KEITT_HOUSE.getUrl(0.42)}
            alt="house keitt"
            class="mango-gallery-image"
            id="smg-keitt-house"
            use:registerAnimatedComponent={keitt_house_animation_frames}
        />
        <img
            src={layout_images.TOMMY_HOUSE.getUrl(0.42)}
            alt="house tommy"
            class="mango-gallery-image"
            id="smg-tommy-house"
            use:registerAnimatedComponent={tommy_house_animation_frames}
        />
        <img
            src={layout_images.DEVELOPMENT_HOUSE_ONE.getUrl(0.42)}
            alt="house in development"
            class="mango-gallery-image"
            id="smg-development-house-one"
            use:registerAnimatedComponent={development_house_one_animation_frames}
        />
        {#if gallery_scroll_progress > 0.54}
            <p id="smg-gallery-caption" transition:fade={{ duration: 700, delay: 200 }}>
                The word Máaku means MANGO in the Wixárika language that is
                spoken by the original towns in the area.
            </p>
        {/if}
    {/if}
</div>

<style>
    #smk-mango-gallery {
        display: grid;
        height: 100%;
        width: 100%;
        grid-template: repeat(21, 1fr) / repeat(21, 1fr);
    }

    :global(#smk-mango-gallery .libery-animation-component) {
        opacity: 0;
        transition: all 1.2s ease;
        z-index: var(--z-index-1);
    }

    .mango-gallery-image {
        width: 100%;
        height: 100%;
        object-fit: cover;
        object-position: center;
    }

    #smg-keitt-house {
        transform: translate(50%, 40%);
        grid-column: 1 / span 8;
        grid-row: 1 / span 15;
    }

    #smg-tommy-house {
        transform: translate(-50%, 40%);
        grid-column: 13 / -1;
        grid-row: 1 / span 11;
    }

    #smg-gallery-caption {
        grid-column: 1 / span 9;
        grid-row: 18 / -1;
    }

    #smg-development-house-one {
        transform: translate(-50%, -40%);
        grid-column: 13 / -1;
        grid-row: 14 / -1;
    }


</style>