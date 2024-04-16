<script>
    import { browser } from "$app/environment";
    import { getTrianglePoints } from "@libs/libery_trigonometry/utils";
    import { createEventDispatcher, onDestroy, onMount } from "svelte";
    import { writable } from "svelte/store";

    
    /*=============================================
    =            Properties            =
    =============================================*/

        /**
         * The video element the component will be controlling
         * @type {HTMLVideoElement}
         */
        export let video_element;

        /**
         * Whether the video is playing or not
         * @type {import('svelte/store').Writable<boolean>}
         */
        export let is_playing = writable(video_element.paused);

        /**
         * Whether the video is muted or not
         * @type {boolean}
         */
        let video_muted = video_element.muted;

        /**
         * The video progress
         * @type {number}     
         */
        let video_progress = 0;

        /**
         * progress indicator drag state
         * @type {boolean}
         */
        let progress_dragging = false;

        /**
         * progress indicator drag start point x
         * @type {number}
         */
        let progress_drag_start_x = 0;

        /**
         * vide progress bar
         * @type {SVGElement}
         */
        let progress_bar;
    
        
        /*----------  Styling  ----------*/

            /**
             * The background color of the controller
             */
            export let controller_background_color = "rgba(0, 0, 0, 0.5)";

            /**
             * Theme color
             * @type {string}
             */
            export let theme_color = "#ff9b55";
        
        const dispatch = createEventDispatcher();    
    
    /*=====  End of Properties  ======*/
    
    onMount(() => {
        video_element.addEventListener('timeupdate', updateVideoProgress);
    })

    onDestroy(() => {
        if (browser) {
            video_element.removeEventListener('timeupdate', updateVideoProgress);
        }
    })

    
    /*=============================================
    =            Methods            =
    =============================================*/

        /**
         * Handles the mute video event
         */
        const handleMuteVideo = () => {
            video_muted = !video_muted;
            video_element.muted = video_muted;
        }

        /**
         * Handles the drag start event on the progress indicator
         * @param {MouseEvent} event
         */
        const handleProgressDragStart = event => {
            event.preventDefault();
            progress_dragging = true;
            progress_drag_start_x = event.clientX; // TODO: this is probably uneeded. it it is, remove it
        }

        /**
         * Handles the drag event on the progress indicator
         * @param {MouseEvent} event
         */
        const handleProgressDrag = e => {
            if (!progress_dragging) return;

            const rect = progress_bar.getBoundingClientRect();
            let client_x_clamp = Math.max(rect.left, Math.min(e.clientX, rect.right));
            video_progress = ((client_x_clamp - rect.left) / rect.width) * 100;    
        }

        /**
         * Handles the drag end event on the progress indicator
         * @param {MouseEvent} event
         */
        const handleProgressDragEnd = event => {
            if (!progress_dragging) return;

            event.preventDefault();
            progress_dragging = false;

            video_element.currentTime = (video_progress / 100) * video_element.duration;
        }
    
        const handlePlayStateChange = () => {
            is_playing.set(video_element.paused);
        }

        /**
         * Handles the click event on the progress bar
         * @param {MouseEvent} event
         */
        const handleProgressClick = event => {
            event.stopPropagation();
            const rect = event.currentTarget.getBoundingClientRect();
            const click_x = event.clientX - rect.left;
            const new_progress = (click_x / rect.width);
            let new_video_time = new_progress * video_element.duration;

            video_element.currentTime = new_video_time;
        }

        const updateVideoProgress = () => {
            if (progress_dragging) return;

            video_progress = (video_element.currentTime / video_element.duration) * 100;
        }
    
    /*=====  End of Methods  ======*/
    
    
    
</script>

<div class="libery-vp-controller"
    style:--player-theme-color={theme_color}
    style:background-color={controller_background_color}
    class:adebug={false}
>
    <div class="libery-vpc-controls">
        <div class="libery-vpc-play-btn-wrapper">
            <svg on:click={handlePlayStateChange} class="libery-vpc-play-state-btn" viewBox="0 0 50 50" role="button" tabindex="0">
                {#if $is_playing}
                    <path d="M10 10L10 40L20 40L20 10ZM30 10L30 40L40 40L40 10L30 10"/>
                {:else}
                    <path d="{getTrianglePoints(25, 25, 25, 0)}"/>
                {/if}
            </svg>
        </div>
        <div class="libery-vpc-video-progress-wrapper" on:mousemove={handleProgressDrag} on:mouseup={handleProgressDragEnd} on:mouseleave={handleProgressDragEnd}>
            <div class="lvc-progress-bar-track">
                <svg bind:this={progress_bar} class="lvc-pbt-track-bar" viewBox="0 0 102 5" on:click={handleProgressClick} >
                    <path class="lvc-pbt-tb-empty-track" d="M 2 2.5 L 100 2.5" />
                    <path class="lvc-pbt-tb-progress-track" d="M 2 2.5 L {video_progress + 2} 2.5" />
                    <circle 
                        on:mousedown={handleProgressDragStart}
                        class="lvc-pbt-tb-progress-indicator"
                        style:transition="all {progress_dragging ? 0 : 0.35}s ease-in-out"
                        cx="{video_progress + 3}"
                        cy="2.5"
                        r="1.5"
                    />
                </svg>
            </div>
        </div>
        <div class="libery-vpc-side-controls">
            <svg on:click={handleMuteVideo} class="libery-vpc-mute-button" role="button" viewBox="0 0 24 24" tabindex="0">
                <path class="outline-path thin" d="M1 18L1 6H5l5 -4V22l-5 -4Z"/>
                {#if !video_muted}
                    <path class="outline-path thin" d="M14 6Q 20 12 14 18"/>
                    <path class="outline-path thin" d="M14 2C 24 2 24 20 14 22"/>
                {:else}
                    <path class="outline-path" d="M14 6L23 18M14 18L23 6"/>
                {/if}
            </svg>
        </div>
    </div>
</div>

<style>
    .libery-vp-controller {
        justify-self: stretch;
        align-self: stretch;
        width: 100%;
        height: 100%;
        container-type: size;
    }

    .libery-vpc-controls {
        display: grid;
        width: 100%;
        height: 100%;
        grid-template-columns: 20% 50% 20%;
        align-items: center;
        gap: 3.3%;
    }
    
    
    /*=============================================
    =            Play state control            =
    =============================================*/
    
     .libery-vpc-play-btn-wrapper {
        display: flex;
        justify-content: flex-end;
        align-items: center;
     }

     svg.libery-vpc-play-state-btn {
        width: 90cqh;
        height: 90cqh;
        fill: var(--player-theme-color);
        outline: none;
     }
    
    /*=====  End of Play state control  ======*/

    
    /*=============================================
    =            Progress bar            =
    =============================================*/

        .libery-vpc-video-progress-wrapper {
            width: 100%;
            height: 100%;
        }
        
        .lvc-progress-bar-track {
            width: 100%;
            height: 100%;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        
        svg.lvc-pbt-track-bar {
            width: 100%;
            height: 100%;
            overflow: visible;
        }

        path.lvc-pbt-tb-empty-track {
            stroke: hsl(from var(--player-theme-color) h 0% l);
            stroke-width: 1px;
            stroke-linecap: round;
        }

        circle.lvc-pbt-tb-progress-indicator {
            fill: var(--player-theme-color);
            transform-box: fill-box;
            transform-origin: center center;
        }

        circle.lvc-pbt-tb-progress-indicator:hover {
            transform: scale(1.8);
            opacity: .5;
        }

        path.lvc-pbt-tb-progress-track {
            stroke: hsl(from var(--player-theme-color) h s calc(l * 0.9));
            stroke-width: 1px;
            stroke-linecap: round;
        }
    
    /*=====  End of Progress bar  ======*/
    
    
    /*=============================================
    =            Side controls            =
    =============================================*/
    
        .libery-vpc-side-controls {
            display: flex;
            justify-content: flex-start;
            align-items: center;
            gap: 10cqh;
        }

        .libery-vpc-side-controls svg{
            width: 90cqh;
            height: 90cqh;
            fill: none;
            outline: none;
        }

        path.outline-path {
            stroke: var(--player-theme-color);
            stroke-width: 1px;
            stroke-linecap: round;
            fill: none;
        }

        path.outline-path.thin {
            stroke-width: 0.5px;
        }

        svg.libery-vpc-mute-button {
            scale: 1 0.7 1;
        }

    
    /*=====  End of Side controls  ======*/
    
    
    
    
</style>