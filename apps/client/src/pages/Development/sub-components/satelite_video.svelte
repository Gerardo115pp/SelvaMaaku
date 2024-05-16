<script>
    import { layout_images, layout_properties } from "@stores/layout";
    import viewport from "@components/viewport_actions/useViewportActions";
    import { fade } from "svelte/transition";
    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        const VIDEO_URL = "/resources/videos/selvamaaku_satelite.webm";

        let video_loaded = false;

        let video_duration_ms = 0;

        /**
         * Whether the section containing the video is visible on the user's viewport
         * @type {boolean}
         */
        let section_visible = false;        

        /**
         * Video shutdown cancelation number
         * @type {number}
         */
        let video_shutdown_cancelation;
    
    /*=====  End of Properties  ======*/

    
    /*=============================================
    =            Methods            =
    =============================================*/


        /**
         * Loads the video
         * @returns {Promise<HTMLVideoElement>}
        */
        const loadVideo = async () => {
            /** @type {HTMLVideoElement} */
            const video_element_loader = document.createElement('video');
    
            let load_promise = new Promise((resolve, reject) => {
                video_element_loader.onloadeddata = () => {
                    video_loaded = true;
                    resolve(video_element_loader);
                }
    
                video_element_loader.onerror = () => {
                    reject();
                } 
            });

            video_element_loader.src = VIDEO_URL;

            return load_promise;
        }
    
        /**
         * Toggles the video play state of the video
        */
        const playVideo = async () => {
            section_visible = true;

            if (!video_loaded) {
                let video_element = await loadVideo();

                video_duration_ms = video_element.duration * 1000;                
            }

            if (video_shutdown_cancelation != null) {
                clearTimeout(video_shutdown_cancelation);
            }
            
            video_shutdown_cancelation = setTimeout(() => {
                shutDownVideo();
                video_shutdown_cancelation = undefined;
            }, video_duration_ms * 0.90);
        }

        /**
         * Shuts down the video
        */
        const shutDownVideo = () => {
            console.log("Shutting down video");
            section_visible = false;
        }
    
    /*=====  End of Methods  ======*/
    
    
    
</script>

<div 
    id="smk-dplss-saylita-map-wrapper"
    class:video-loading={!video_loaded && section_visible} 
    on:viewportEnter={playVideo}
    on:viewportLeave={shutDownVideo}
    use:viewport={{height_offset: $layout_properties.IS_MOBILE ? 0.9 : 0.5}}
>
    {#if video_loaded && section_visible}
        <video src={VIDEO_URL} autoplay muted playsinline></video>
    {:else}
         <img class:video-loaded={video_loaded} out:fade={{ duration: 600 }} src="{layout_images.SAYULITA_MAP.getUrl(1)}" alt="sayulita map">
    {/if}
</div>

<style>
    #smk-dplss-saylita-map-wrapper {
        --default-component-height: 740px;

        position: relative;
        width: 100%;
        max-height: var(--default-component-height);
        height: var(--default-component-height);
        overflow: hidden;
    }

    /* #smk-dplss-saylita-map-wrapper:empty {
    } */

    #smk-dplss-saylita-map-wrapper img, #smk-dplss-saylita-map-wrapper video {
        width: 100%;
        max-height: var(--default-component-height);
        object-fit: cover;
        object-position: left bottom;
        pointer-events: none;
    }

    #smk-dplss-saylita-map-wrapper > video {
        object-position: center;
    }

    #smk-dplss-saylita-map-wrapper img.video-loaded {
        height: var(--default-component-height);
        position: absolute;
        top: 0;
        left: 0;
    }

    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
    @media only screen and  (max-width: 768px) {
        #smk-dplss-saylita-map-wrapper {
            --default-component-height: 216px;
        }
    }
    
    /*=====  End of Mobile  ======*/
    
    
</style>