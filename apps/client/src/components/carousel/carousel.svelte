<script>
    import { onDestroy, onMount } from "svelte";
    import { browser } from "$app/environment";
    import { CarouselItem } from "./types";
    import { writable, get } from "svelte/store";

    /*=============================================
    =            Properties            =
    =============================================*/

         /**
         * Array of objects that contain the images of the carousel
         * @type {CarouselItem[]}
         */
        export let carousel_items = [];
        
        /*----------  Behavior props  ----------*/
                

            /**
             * Index of the current slide
             * @type {import("svelte/store").Writable<number>}
             */
            export let current_slide_index = writable(0);

            /**
             * Whether to loop the carousel
             * @type {boolean}
             * @default false
             */
            export let loop_slides = false;

            /**
             * The amount of slides that will be shown at once. If not set then the CarouselItem.width will be used
             * but if set, the CarouselItem.width will be ignored
             * @type {number}
             */
            export let visible_slides;

            /**
             * Whether or not the autoplay is enabled
             * @type {boolean}
             * @default false
             */
            export let autoplay = false;

            /**
             * The interval in milliseconds between each slide
             * @type {number}
             * @default 1000
             */
            export let autoplay_interval_ms = 1000;

            export let is_marquee = false;
            $: if (is_marquee) {
                autoplay = false;
                autoplay_interval_ms = 0;
            }

        /*----------  End of Behavior props  ----------*/

        /**
         * The interval object that will be used to autoplay the carousel. 
         * Use it to clear the interval when the component will be destroyed
         * @type {number}
         */
        let autoplay_interval;

        
        /*----------  Styling  ----------*/
        
            /**
             * The height of the slides
             * @type {string}
             * @default "333px"
             */
            export let slide_height = "333px";

            /**
             * Whether to use a gradient mask to hide the edges of the carousel
             * @type {boolean}
            */
            export let use_mask = false;

            /**
             * The gap between the slides
             * @type {string}
             * @default "var(--spacing-2)"
            */
            export let carousel_gap = "var(--spacing-2)";
        

    /*=====  End of Properties  ======*/ 

    onMount(() => {
        if (autoplay && !is_marquee) {
            autoplay_interval = window.setInterval(handleAutoplay, autoplay_interval_ms);
        }
    });

    onDestroy(() => {
        if (browser && autoplay_interval != null) {
            window.clearInterval(autoplay_interval);
        }
    });

    
    /*=============================================
    =            Methods            =
    =============================================*/

        /**
         * Handles the change of the slide
         * @param {number} step The amount of steps to move the slide. Default is 1
         */
        const changeSlide = (step=1) => {
            console.log('current_slide_index:', $current_slide_index);
            let new_slide_index = $current_slide_index + step;   
            console.log('new_slide_index:', new_slide_index);

            if (loop_slides) {
                new_slide_index = new_slide_index % carousel_items.length;
            }

            current_slide_index.set(Math.max(0, Math.min(new_slide_index, carousel_items.length - 1)));
            console.log('current_slide_index:', $current_slide_index);
        }

        const handleAutoplay = () => {
            // Check if the autoplay is not enabled. if it is not, clear the interval(if we are in the browser)
            if (!autoplay) {
                if (browser) {
                    window.clearInterval(autoplay_interval);
                }

                return;
            }

            changeSlide();
        }
    
    /*=====  End of Methods  ======*/
    
    
</script>


{#if carousel_items.length > 0}
    <div class="libery-carousel-component" 
        class:lcc-marquee-carousel={is_marquee}
        style:--carousel-gap={carousel_gap}
        style:--slide-width="{100 / visible_slides}cqw"
    >
        {#if $$slots.arrow_left}
            <div class="carousel-arrow carousel-arrow-left" on:click={() => changeSlide(-1)}>
                <slot name="arrow_left" />
            </div>
        {/if}
        <div class="lcc-carousel-mask"
            class:gradient-mask={use_mask}
             class:dtwo={false}
        >
            <ol class="lcc-carousel-slide-stack"
                style:translate="calc({-1 * $current_slide_index} * calc(var(--slide-width) + var(--carousel-gap)))"
            >
                {#each carousel_items as slide, index}
                    {@const is_slide_active = index === $current_slide_index}
                    <li class="lcc-carousel-slide"
                        style:height="{slide_height}"
                        style:width='var(--slide-width)'
                        class:lcc-carousel-slide--active={is_slide_active}
                        class:dthree={false}
                    >
                        <img src={slide.src} alt="Slide {index + 1}" />
                    </li>
                {/each}
                {#if is_marquee}
                    {#each carousel_items as slide, index}
                        <li class="lcc-carousel-slide lcc-carousel-slide--doppelganger"
                            style:height="{slide_height}"
                            style:width={visible_slides ? `${100 / visible_slides}cqw` : `${slide.width}px`}    
                            class:dfour={false}
                            aria-hidden="true"
                        >
                            <img src={slide.src} alt="Slide {index + 1}" />
                        </li>
                    {/each}
                {/if}
            </ol>
        </div>
        {#if $$slots.arrow_right}
            <div class="carousel-arrow carousel-arrow-right" on:click={() => changeSlide()}>
                <slot name="arrow_right" />
            </div>
        {/if}
    </div>
{/if}

<style>
    .libery-carousel-component {
        width: 100%;
        container: inline-size;
        padding: 0 var(--spacing-3);
    }

    .libery-carousel-component:has(.carousel-arrow) {
        display: grid;
        grid-template-columns: auto 1fr auto;
        align-items: center;
        gap: var(--spacing-2);
    }

    .lcc-carousel-mask {
        overflow: hidden;
        container-type: inline-size;
        width: 100%;
        max-width: 100cqw;
        padding: var(--spacing-3) 0;
    }
    
    .lcc-carousel-mask.gradient-mask {
        mask: linear-gradient(90deg, #ffffff01 2%, #ffffff8c 30%, white 90%, transparent 100%);
    }

    /* use to debug the carousel */
    .lcc-carousel-mask.dtwo {
        overflow: visible;
    }

    @keyframes marquee-slide-stack {
        to {
            transform: translateX(calc(-50% - calc(var(--carousel-gap) / 2)));
        }
    }

    ol.lcc-carousel-slide-stack {
        display: flex;
        flex-direction: row;
        width: fit-content;
        flex-wrap: nowrap;
        list-style: none;
        gap: var(--carousel-gap);
        margin: 0;
        padding: 0;
        transition: translate 0.5s ease-in-out;
    }

    .lcc-marquee-carousel ol.lcc-carousel-slide-stack {
        animation: marquee-slide-stack 60s linear infinite;
        transition: none;
    }

    .lcc-carousel-slide {
        flex-shrink: 0;
        box-shadow: 2px 7px 9px 2px rgba(0, 0, 0, 0.15);
        user-select: none;

        & img {
            width: 100%;
            height: 100%;
            object-fit: cover;
            pointer-events: none;
        }
    }
</style>