<script>
    import { getTrianglePoints } from "@libs/libery_trigonometry/utils";
    import { writable } from "svelte/store";
    
    /*=============================================
    =            Properties            =
    =============================================*/

        /** 
         * Whether the video is playing or not
         * @type {import('svelte/store').Writable<boolean>}
         */
        export let is_playing = writable(false);

        /*----------  Styling  ----------*/
        
            /**
             * Whether to add a background to the button
             * @type {boolean}
             * @default true
             */
            export let has_background = true;

            /**
             * The size of the button
             * @type {string}
             * @default "10cqw"
             */
            export let size = "10cqw";
        
    
    /*=====  End of Properties  ======*/
    
    
    /*=============================================
    =            Methods            =
    =============================================*/
    
        const handlePlayStateChange = () => {
            is_playing.set(!$is_playing);
        }
    
    /*=====  End of Methods  ======*/
    
    
</script>

<svg 
    on:click={handlePlayStateChange}
    class="play-state-controller-btn"
    role="button"
    viewBox="0 0 50 50"
    tabindex="0"
    style:width={size}
    style:height={size}
    class:is-playing={$is_playing}
    class:backgroundless={!has_background}
>
    <circle cx="25" cy="25" r="25"/>
    {#if !$is_playing}
         <path d="{getTrianglePoints(25, 25, 25, 0)}"/>
    {:else}
        <path d="M10 10L10 40L20 40L20 10ZM30 10L30 40L40 40L40 10L30 10"/>
    {/if}
</svg>

<style>

    .play-state-controller-btn {
        border-radius: 100%;
        backdrop-filter: blur(5px);
        overflow: visible;
        outline: hsla(0, 100%, 100%, 0.2) solid 0.1px;
        
        & circle {
            fill: hsl(from var(--theme-color-light-active) h s l / 0.2);
        }

        & path {
            fill: hsl(from var(--theme-color) h s l / 0.9);
            transition: fill 0.3s ease-in-out;
            pointer-events: none;
        }

        &:has(circle:hover) path {
            fill: hsl(from var(--theme-color) h calc(s * 1.1) calc(l * 0.9) / 1);
        }
    }

    .is-playing.play-state-controller-btn {
        transition: opacity 0.7s ease-in-out;

        & path {
            fill: hsl(from var(--theme-color) h s l / 0.9);
            transition: fill 0.3s ease-in-out;
            pointer-events: none;
            transform-box: fill-box;
            transform-origin: center;
            scale: 0.9 1 1;
        }
    }

    .backgroundless.play-state-controller-btn {
        backdrop-filter: none;
        background: transparent;
        outline: none;

        & circle {
            fill: transparent;
        }

        & path {
            fill: hsl(from var(--theme-color) h calc(s * 1.2) 40% / 1);
        }
    }

    @supports not (color: rgb(from white r g b)) {
        

        .play-state-controller-btn circle {
            fill: #b7bfbc42;
        }

        .play-state-controller-btn path {
            fill: color-mix(in srgb, var(--theme-color), rgba(255, 255, 255, 0.39) 10%);
            opacity: 0.8;
        }

        .play-state-controller-btn:has(circle:hover) path {
            fill: color-mix(in srgb, var(--theme-color), green 10%);
        }
    }
</style>