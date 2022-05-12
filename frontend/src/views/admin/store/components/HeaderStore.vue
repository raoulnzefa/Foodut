<template>
    <vx-card>
        <!-- swiper -->
        <swiper :options="swiperOption" :dir="$vs.rtl ? 'rtl' : 'ltr'" :key="$vs.rtl">
            <div class="parallax-bg" slot="parallax-bg" data-swiper-parallax="-23%"></div>
            <swiper-slide>
                <div class="title" data-swiper-parallax="-100">{{ storeName }}</div>
                <div class="subtitle font-semibold" data-swiper-parallax="-200">{{ city }}</div>
                <div class="text" data-swiper-parallax="-300">
                    <p class="font-medium">Jelly chocolate cupcake chocolate bar caramels chupa chups chupa chups ice cream tiramisu. Oat cake muffin pastry marzipan sweet jujubes powder cupcake carrot cake. Caramels candy pastry marzipan pudding. Sugar plum carrot cake topping cookie.</p>
                </div>
            </swiper-slide>
            <swiper-slide>
                <div class="title" data-swiper-parallax="-100">Slide 2</div>
                <div class="subtitle font-semibold" data-swiper-parallax="-200">Subtitle</div>
                <div class="text" data-swiper-parallax="-300">
                    <p class="font-medium">Sweet oat cake marzipan jelly brownie ice cream bear claw marshmallow jelly beans. Halvah dessert caramels toffee sweet cake tiramisu. Chocolate bar marshmallow topping biscuit gummies chocolate cake candy liquorice.</p>
                </div>
            </swiper-slide>
            <swiper-slide>
                <div class="title" data-swiper-parallax="-100">Slide 3</div>
                <div class="subtitle font-semibold" data-swiper-parallax="-200">Subtitle</div>
                <div class="text" data-swiper-parallax="-300">
                    <p class="font-medium">Powder jelly bonbon liquorice candy canes jujubes fruitcake cotton candy macaroon. Croissant jelly-o tootsie roll halvah. Topping lollipop pastry. Fruitcake powder cupcake apple pie chocolate bar wafer. brownie macaroon dragée chocolate bar candy.</p>
                </div>
            </swiper-slide>
            <div class="swiper-pagination swiper-pagination-white" slot="pagination"></div>
            <div class="swiper-button-prev swiper-button-white" slot="button-prev"></div>
            <div class="swiper-button-next swiper-button-white" slot="button-next"></div>
        </swiper>
    </vx-card>
</template>

<script>
import 'swiper/dist/css/swiper.min.css'
import { swiper, swiperSlide } from 'vue-awesome-swiper'

export default {
  props: {
    storeName: String,
    city: String
  },
  data () {
    return {
      swiperOption: {
        speed: 600,
        parallax: true,
        pagination: {
          el: '.swiper-pagination',
          clickable: true
        },
        navigation: {
          nextEl: '.swiper-button-next',
          prevEl: '.swiper-button-prev'
        }
      }
    }
  },
  components: {
    swiper,
    swiperSlide
  }
}
</script>

<style scoped>
.swiper-slide {
    font-size: 18px;
    color: #fff;
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
    padding: 40px 60px;
    background-color: transparent !important;
    justify-content: space-around !important;
}

.parallax-bg {
    position: absolute;
    left: 0;
    top: 0;
    width: 130%;
    height: 100%;
    -webkit-background-size: cover;
    background-size: cover;
    background-position: center;
    background-image: url('../../../../assets/images/pages/carousel/banner-4.jpg');
}

.swiper-slide .title {
    font-size: 41px;
    font-weight: 300;
}

.swiper-slide .subtitle {
    font-size: 21px;
}

.swiper-slide .text {
    font-size: 14px;
    max-width: 400px;
    line-height: 1.3;
}

</style>
