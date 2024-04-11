<script>
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import { layout_images } from "@stores/layout";
</script>

<section id="smk-hmp-cocoteros-slides">
    <div id="smk-hmpcs-content-wrapper" class="design-content-width">
        <div class="wixarica-icons">
            <WixaricaIcon with_accent />
        </div>
        <div id="smk-hmpcs-slides-wrapper">
            <!-- TODO: Make the actual carousel -->
            <h2 id="smk-hmpcs-carousel-title">
                Cocoteros House Model
            </h2>
            <div class="smk-fake-carousel">
                <div class="smk-fake-carousel-arrow-left smk-fake-carousel-arrow">
                    <svg width="48" height="48" viewBox="0 0 48 48">
                        <rect width="48" height="48" rx="24" fill="#B7BFBC"/>
                        <path d="M31 24H17M17 24L24 31M17 24L24 17" stroke="#183028" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                    </svg>    
                </div>
                <div class="smk-fake-carousel-track">
                    <img src="{layout_images.COCOTEROS_HOUSE_TWO.getUrl(0.8)}" alt="Cocoteros House Model" />
                </div>
                <div class="smk-fake-carousel-arrow-right smk-fake-carousel-arrow">
                    <svg width="48" height="48" viewBox="0 0 48 48">
                        <rect width="48" height="48" rx="24" fill="#B7BFBC"/>
                        <path d="M31 24H17M17 24L24 31M17 24L24 17" stroke="#183028" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                    </svg>    
                </div>
            </div>
        </div>
    </div>
</section>

<style>
    #smk-hmp-cocoteros-slides {
        padding: var(--spacing-5) 0;
    }

    #smk-hmpcs-content-wrapper {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: var(--spacing-4);
    }

    
    /*=============================================
    =            FakeCarousel            =
    =============================================*/
    
        #smk-hmpcs-slides-wrapper {
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: var(--spacing-3);
        }

        .smk-fake-carousel {
            display: flex;
            justify-content: space-between;
            align-items: center;
        }

        .smk-fake-carousel-track {
            width: min(82.73026%, 1006px);
        }

        .smk-fake-carousel-arrow-right {
            rotate: 180deg;
        }
    
    /*=====  End of FakeCarousel  ======*/
    
    
</style>