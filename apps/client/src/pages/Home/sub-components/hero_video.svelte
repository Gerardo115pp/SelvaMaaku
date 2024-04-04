<script>
    import VideoPlayer from "@components/VideoPlayer/VideoPlayer.svelte";
    import { layout_images } from "@stores/layout";

    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        //* Style properties
        /**
         * The width of the video player.
         * @type {string}
         * @default "max-content"
         */
        export let video_width = "max-content";
        /**
         * The height of the video player.
         * @type {string}
         * @default "auto"
         */
        export let video_height = "auto";
    
    /*=====  End of Properties  ======*/
    
    
</script>

<section id="smk-hero-video">
    <VideoPlayer 
        poster_media={layout_images.HERO_VIDEO_COVER}
        poster_threshold={0.8}
        player_width={video_width}
        player_height={video_height}
    />
</section>

<style>
    section#smk-hero-video {
        width: max-content;
    }
</style>