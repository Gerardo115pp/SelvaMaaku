<script>
    import InteractiveImage from "@components/Images/InteractiveImage.svelte";
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import { layout_properties, layout_images } from "@stores/layout";
    import { setRenderMedia } from "@stores/media_display";
    import { elasticOut } from "svelte/easing";
    import { fly } from "svelte/transition";

    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        /**
         * Whether the gallery is visible
         * @type {boolean}
         */
        export let gallery_visible = false;

        
        /*----------  Entry animation  ----------*/
        
        

            /*
            * The duration of the entry animation in milliseconds
            * @type {number}
            */
            export let entry_animation_duration = 1500;

            const fly_right_transition = {
                delay: 200,
                duration: entry_animation_duration * 0.95,
                easing: elasticOut,
                x: 500,
                y: 120,
                opacity: 0,
            }

            const fly_left_transition = {
                delay: fly_right_transition.delay,
                duration: fly_right_transition.duration,
                easing: fly_right_transition.easing,
                x: fly_right_transition.x * -1,
                y: fly_right_transition.y * -1,
                opacity: fly_right_transition.opacity,
            }
        
        
        /*----------  Mouse tracking animation  ----------*/

            let mouse_x = 0;
            let mouse_y = 0;

            let mouse_inside_component = false;

            const mouse_tracking_factor = 17;

            const position_difference_threshold = 0.1;

            /**
             * The particle that are going to be tracking the mouse
             * @type {HTMLElement[]}
            */
            let mouse_tracking_particles = [];

            /**
             * Gallery component
             * @type {HTMLElement}
             */
            let gallery_component;
            
    /*=====  End of Properties  ======*/

    
    /*=============================================
    =            Methods            =
    =============================================*/

            const handleMouseEnter = () => {
                console.log('mouse enter');
                mouse_inside_component = true;

                // for (let particle of mouse_tracking_particles) {
                //     particle.style.transition = 'transform 0.05s';
                // }
            }

            const handleMouseLeave = () => {
                console.log('mouse leave');
                mouse_inside_component = false;

                for (let particle of mouse_tracking_particles) {
                    // particle.style.transition = 'transform 1s';
                    particle.style.transform = 'translate(0, 0)';
                }
            }
    
            /**
             * Registers the mouse tracking particle    
             * @param {HTMLElement} node
             */
            const registersMouseTrackingParticle = node => {
                if (!(node instanceof HTMLElement) || $layout_properties.IS_MOBILE) return;

                mouse_tracking_particles.push(node);

                return {
                    destroy() {
                        if (mouse_tracking_particles?.length > 0) {
                            mouse_tracking_particles = mouse_tracking_particles.filter(particle => particle !== node);
                        }
                    }
                }
            }

            /**
             * Updates the mouse position
             * @param {MouseEvent} e
             */
            const updateMousePosition = e => {
                if (gallery_component == null || $layout_properties.IS_MOBILE) return;

                const component_rect = gallery_component.getBoundingClientRect();

                let x_in_px = e.clientX - component_rect.left;
                let y_in_px = e.clientY - component_rect.top;

                let new_mouse_x = (x_in_px / component_rect.width) * mouse_tracking_factor;
                let new_mouse_y = (y_in_px / component_rect.height) * mouse_tracking_factor;

                let euclidean_distance = Math.sqrt((new_mouse_x - mouse_x) ** 2 + (new_mouse_y - mouse_y) ** 2);

                if (euclidean_distance >= position_difference_threshold) {
                    mouse_x = new_mouse_x;
                    mouse_y = new_mouse_y;
                }
                    

                // console.log(`mouse position: (${mouse_x}, ${mouse_y})`);
            }
    
    /*=====  End of Methods  ======*/
    
    
    
    

</script>

<div 
    id="smk-models-gallery"
    bind:this={gallery_component}
    on:mousemove={updateMousePosition}
    on:mouseenter={handleMouseEnter}
    on:mouseleave={handleMouseLeave}
    class:is-visible={gallery_visible}
    class:is-mouse-inside={mouse_inside_component}
>
    {#if gallery_visible}
         <!-- content here -->
        <div id="smk-mg-wixarica-top"
            class="wixarica-gallery-icon tracking-particle"
            style="transform: translate({mouse_x * 2.3}%, {(mouse_y * 1.5)}%);"
            in:fly={{
                ...fly_right_transition,
                delay: fly_right_transition.delay * 1.45,
            }}
            use:registersMouseTrackingParticle
        >
            <WixaricaIcon opacity={0.5} is_flat_color />
        </div>
        <div id="smk-mg-wixarica-bottom"
            class="wixarica-gallery-icon tracking-particle"
            style="transform: translate({mouse_x * 2.3}%, {-(mouse_y * 1.5)}%);"
            in:fly={{
                ...fly_left_transition,
                delay: fly_left_transition.delay * 1.5,
            }}
            use:registersMouseTrackingParticle
        >
            <WixaricaIcon opacity={0.5} is_flat_color />
        </div>
        <!-- TODO: Add an alt text to the images -->
        <div id="smk-mg-keitt-house" 
            class="smk-mg-gallery-image tracking-particle"
            style="transform: translate({-(mouse_x * 0.1)}%, {-(mouse_y * 0.1)}%);"
            in:fly={{
                ...fly_left_transition,
                delay: fly_left_transition.delay * 2,
            }}
            use:registersMouseTrackingParticle
        >
            <InteractiveImage 
                on:image-clicked={() => setRenderMedia(layout_images.KEITT_HOUSE.getUrl(1))}
                image_resource={layout_images.KEITT_HOUSE}
                desktop_viewport_percentage={0.4}
                mobile_viewport_percentage={1}
                overlay_color="var(--theme-color)"
                max_overlay_opacity={0.5}
            />
        </div>
        <div id="smk-mg-keitt-house-two"
            class="smk-mg-gallery-image tracking-particle"
            style="transform: translate({(mouse_x * 0.1)}%, {-(mouse_y * 0.1)}%);"
            in:fly={{
                ...fly_right_transition,
                delay: fly_right_transition.delay * 2.5,
            }}
            use:registersMouseTrackingParticle
        >
            <InteractiveImage 
                on:image-clicked={() => setRenderMedia(layout_images.KEITT_HOUSE_TWO.getUrl(1))}
                image_resource={layout_images.KEITT_HOUSE_TWO}
                desktop_viewport_percentage={0.4}
                mobile_viewport_percentage={0.8}
                overlay_color="var(--theme-color)"
                max_overlay_opacity={0.5}
            />
        </div>
        <div id="smk-mg-tommy-house-three"
            class="smk-mg-gallery-image tracking-particle"
            style="transform: translate({-(mouse_x * 0.4)}%, {-(mouse_y * 1)}%);"
            in:fly={{
                ...fly_left_transition,
                delay: fly_left_transition.delay * 1.3,
            }}
            use:registersMouseTrackingParticle
        >
            <InteractiveImage 
                on:image-clicked={() => setRenderMedia(layout_images.TOMMY_HOUSE_THREE.getUrl(1))}
                image_resource={layout_images.TOMMY_HOUSE_THREE}
                desktop_viewport_percentage={0.3}
                mobile_viewport_percentage={0.6}
                overlay_color="var(--theme-color)"
                max_overlay_opacity={0.5}
            />
        </div>
        <div id="smk-mg-tommy-house-four"
            class="smk-mg-gallery-image tracking-particle"
            style="transform: translate({(mouse_x * 0.3)}%, {(mouse_y * 0.3)}%);"
            in:fly={{
                ...fly_right_transition,
                delay: fly_right_transition.delay * 1.7,
            }}
            use:registersMouseTrackingParticle
        >
            <InteractiveImage 
                on:image-clicked={() => setRenderMedia(layout_images.TOMMY_HOUSE_FOUR.getUrl(1))}
                image_resource={layout_images.TOMMY_HOUSE_FOUR}
                desktop_viewport_percentage={0.4}
                mobile_viewport_percentage={0.8}
                overlay_color="var(--theme-color)"
                max_overlay_opacity={0.5}
            />
        </div>
        <span id="smk-mg-keitt-label"
            class="gallery-label tracking-particle"
            style="transform: translate({-(mouse_x * 0.5)}%, {-(mouse_y * 0.05)}%);"
            in:fly={fly_right_transition}
            use:registersMouseTrackingParticle
        >
            Keitt
        </span>
        <span id="smk-mg-tommy-label"
            class="gallery-label tracking-particle"
            style="transform: translate({(mouse_x * 0.5)}%, {-(mouse_y * 0.05)}%);"
            in:fly={fly_left_transition}
            use:registersMouseTrackingParticle
        >
            Tommy
        </span>
    {/if}
</div>

<style>
    #smk-models-gallery {
        position: relative;
        width: min(624px, 51.31578cqw);
        height: min(765px, 62.91118cqw);
        container-type: inline-size;
    }

    .smk-mg-gallery-image {
        object-fit: cover;
        object-position: center;
    }

    .wixarica-gallery-icon {
        width: 12.820%;
    }

    .gallery-label {
        font-family: var(--font-titles);
        color: var(--light-green);
        font-size: 15.3846cqw;
        font-weight: 300;
        line-height: 1;
        
        &:hover {
            cursor: default;
            color: hsl(from var(--light-green) h 37% 65%);
            text-shadow: 0 0 5px hsl(from var(--light-green) h 37% 65% / 0.6);
            transition: text-shadow 0.5s ease-in-out, color 0.4s ease-in-out !important;
        }
    }

    
    /*=============================================
    =            Items positions            =
    =============================================*/
    
        #smk-mg-wixarica-top {
            position: absolute;
            left: 74.35%;
            top: 0;
        } 

        #smk-mg-wixarica-bottom {
            position: absolute;
            left: 21.47%;
            bottom: 0;
        }

        #smk-mg-keitt-house {
            position: absolute;
            left: 14.42%;
            top: 0;
            width: 52.24358%;
            height: 65.35947%;
        }

        #smk-mg-keitt-house-two {
            position: absolute;
            left: 0;
            top: 11.37%;
            width: 33.17307%;
            height: 18.69281%;
        }

        #smk-mg-tommy-house-three {
            position: absolute;
            left: 0;
            top: 54.37%;
            width: 33.17307%;
            height: 21.56862%;
        }

        #smk-mg-tommy-house-four {
            position: absolute;
            right: 0;
            bottom: 0;
            width: 44.87179%;
            height: 49.6732%;
        }

        #smk-mg-keitt-label {
            position: absolute;
            left: 63.14%;
            top: 20.65%;
        }

        #smk-mg-tommy-label {
            position: absolute;
            left: 0.8%;
            top: 68.62%;
        }


    
    /*=====  End of Items positions  ======*/
    
    
    /*=============================================
    =            Tracking particles            =
    =============================================*/
    
        .is-mouse-inside .tracking-particle {
            transition: none;
        }

        .tracking-particle {
            transition: transform 1s ease-out;
        }
    
    /*=====  End of Tracking particles  ======*/
    
    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @container (width <= 950px) {
            #smk-models-gallery {
                position: relative;
                width: 100cqw;
                height: 122.59615cqw;
                container-type: inline-size;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
    
</style>