<script>
    import { fade } from 'svelte/transition';
    import VideoController from './VideoController.svelte';
    import { getTrianglePoints } from '@libs/libery_trigonometry/utils';
    import PlayStateManager from './buttons/playstate_manager.svelte';
    import { writable } from 'svelte/store';
    import { onMount } from 'svelte';

    
    /*=============================================
    =            Properties            =
    =============================================*/

        /**
         * The video player element.
         * @type {HTMLElement}
         */
        let video_player;

        /**
         * The actual video.
         * @type {HTMLVideoElement}
         */
        let video_element;

    
        export let video_url = '';

        
        /*----------  Customization  ----------*/
        
            

            /**
             * The resource that will be used as the poster for the video.
             * @type {import('@models/ImageResources').ImageResource}
             * @default null
             */
            export let poster_media;

            /**
             * The percentage of the viewport that the video poster will be loaded for.
             * @type {number}
             * @default 0.5
            */
            export let poster_threshold = 0.5;

            //* Style properties
            
            /**
             * The width of the video player.
             * @type {string}
             * @default "max-content"
             */
            export let player_width = "max-content";
            /**
             * The height of the video player.
             * @type {string}
             * @default "auto"
             */
            export let player_height = "auto";

            /**
             * The alt text 
            */
            export let alt_text = 'the video player component';

        
        /*----------  State  ----------*/
        
            /**
             * The player reproduction state
             * @type {import('svelte/store').Writable<boolean>}
             */
            let is_playing = writable(false);

            /**
             * The video load state
             * @type {boolean}
             */
            let is_loaded = false;
    
    /*=====  End of Properties  ======*/

    onMount(() => {
        is_playing.subscribe(() => {
            togglePlayState();
        });
    })

    
    /*=============================================
    =            Methods            =
    =============================================*/

        const loadVideo = async () => {
            /**
             * @type {HTMLVideoElement} 
             */
            const video_element_loader = document.createElement('video');

            let load_promise = new Promise((resolve, reject) => {
                video_element_loader.onloadeddata = () => {
                    is_loaded = true;
                    resolve();
                }

                video_element_loader.onerror = () => {
                    reject();
                }
            });

            video_element_loader.src = video_url;

            return load_promise;
        }

        const togglePlayState = () => {
            if (video_player == null) return;

            if (!is_loaded && video_url !== '' && $is_playing) {
                loadVideo();
            } else if (is_loaded && video_element != null) {
                if ($is_playing) {
                    video_element.play();
                } else {
                    video_element.pause();
                }
            }
        }
    
    /*=====  End of Methods  ======*/
    
    
    
    
</script>

<div 
    class="video-player-component"
    style:width={player_width}
    style:height={player_height}
    class:is-playing={$is_playing}
    class:video-size-fixed={!(player_height === 'auto' || player_width === 'auto')}
    class:is-loaded={is_loaded}
    class:is-loading={!is_loaded && video_url !== '' && $is_playing}
    bind:this={video_player}
>
    <div class="video-player-controls-overlay">
        <div class="video-player-mega-controls">
            <PlayStateManager 
                {is_playing}
            />
        </div>
        {#if !is_loaded && video_url !== '' && $is_playing}
            <div in:fade={{duration: 700}} class="libery-video-loader"></div>
        {/if}
        {#if is_loaded && video_element != null }
            <div class="video-player-controls">
                <VideoController 
                    bind:video_element
                    controller_background_color="hsl(from var(--theme-color-darker) h s l / 0.95)"
                    theme_color="hsl(from var(--theme-color) h calc(s * 1.1) 30%)"
                    {is_playing}
                />
            </div>
        {/if}
    </div>
    {#if video_url === '' || !is_loaded}
        <div class="video-player-poster-wrapper">
            <img src="{poster_media.getUrl(poster_threshold)}" alt="{alt_text}">
        </div>
    {:else if is_loaded}
        <video 
            bind:this={video_element}
            src="{video_url}"
            autoplay
            loop
        >
            <track kind="captions"/>
        </video>
    {/if}
</div>

<style>
    .video-player-component {
        position: relative;
        container-type: size;
    }

    .video-player-component video {
        pointer-events: none;
    }

    .video-player-component.video-size-fixed video {
        width: 100cqw;
        height: 100cqh;
        object-fit: cover;
        object-position: center;
    }


    
    /*=============================================
    =            Controls overlay            =
    =============================================*/
    
        .video-player-controls-overlay {
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            display: grid;
            grid-template: 
                "vps vps vps" auto
                "vpc vpc vpc" 7%
            ;
            justify-items: center;
            align-items: center;
            will-change: background;
            transition: background 0.7s ease-in-out;
        }

        .is-loading .video-player-controls-overlay {
            background: hsl(from var(--theme-color) h calc(s * 1.4) l / 0.6);
        }

        .libery-video-loader {
            --loader-size: 80px;
            grid-area: vps;

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

        .video-player-mega-controls {
            grid-area: vps;
            width: 100%;
            height: 100%;
            display: grid;
            place-items: center;
        }

        .is-playing .video-player-mega-controls {
            opacity: 0;
            transition: opacity 0.7s ease-in-out;
        }

        .is-playing .video-player-mega-controls:hover {
            opacity: 1;
        }

        .video-player-controls {
            grid-area: vpc;
            width: 100%;
            height: 100%;
        }
    
    /*=====  End of Controls overlay  ======*/
    
    

    .video-player-poster-wrapper img {
        width: 100cqw;
        height: 100cqh;
        object-fit: cover;
        object-position: center;
        user-select: none;
        pointer-events: none;
    }
    
    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @container (width <= 750px) {
            .video-player-controls-overlay {
                grid-template: 
                    "vps vps vps" auto
                    "vpc vpc vpc" 10%
                ;            
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
</style>