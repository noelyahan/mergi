package main

import (
	"os"
	"testing"
)

func TestProductsMerge(t *testing.T) {
	args := []string{
		"mergi",
		"-t", "TTT",
		"-i", "https://www.mensjournal.com/wp-content/uploads/mf/1_loreal.jpg",
		"-r", "355 236",
		"-i", "http://cdn.shopify.com/s/files/1/0737/8455/products/0008_39A_Palette_Stylized_7454c378-d527-4655-9898-55ae34103e98_grande.jpg?v=1532967345",
		"-r", "355 236",
		"-i", "https://off.com.ph/-/media/images/off/ph/products-en/products-landing/landing/off_overtime_product_collections_large_2x.jpg",
		"-r", "355 236",
		"-a", "smooth 10",
		"-o",
		"out.gif",
	}
	os.Args = args
	main()
}

func TestFacesMerge(t *testing.T) {
	args := []string{
		"mergi",
		"-t", "TTT",
		"-i", "https://d2v9y0dukr6mq2.cloudfront.net/video/thumbnail/4e_S7GBpeilxvw00j/beautiful-healthy-smiling-woman-with-fresh-skin-of-face-over-white-background_45uvpld6x__S0000.jpg",
		"-r", "300 300",
		"-i", "https://thumbs.dreamstime.com/t/beautiful-young-woman-face-over-white-background-health-people-plastic-surgery-beauty-concept-73966507.jpg",
		"-r", "300 300",
		"-i", "http://s1.1zoom.me/big0/639/White_background_Face_Glance_Beautiful_Blonde_girl_538527_1280x852.jpg",
		"-r", "300 300",
		"-a", "smooth 5",
		"-o",
		"out.gif",
	}
	os.Args = args
	main()
}

func TestFlowersAnim(t *testing.T) {
	args := []string{
		"mergi",
		"-t", "TT",
		"-i", "https://cdn.pixabay.com/photo/2013/09/29/09/52/aster-188045_960_720.jpg",
		"-r", "492 326",
		"-c", "100 80 250 200",
		"-i", "https://cdn.pixabay.com/photo/2015/04/19/08/33/flower-729512_960_720.jpg",
		"-r", "492 326",
		"-c", "100 80 250 200",
		"-a", "smooth 5",
		"-o",
		"out.gif",
	}
	os.Args = args
	main()
}

func TestMergiCommands(t *testing.T) {
	args := []string{
		"mergi",
		"-t",
		"TTTT",
		//"-i",
		//"https://i.ytimg.com/vi/HNAM2EVXH9A/maxresdefault.jpg",
		//"-w",
		//"http://watermarkrestaurant.net/wp-content/uploads/2017/03/cropped-watermark-restaurant.png -r 500 200 700 520",
		//"-w",
		//"http://watermarkrestaurant.net/wp-content/uploads/2017/03/cropped-watermark-restaurant.png -r 500 200 0 0",

		//"-c",
		//"0 0 160 160",
		//"-r",
		//"320 320",
		"-i",
		"https://i.stack.imgur.com/Ni7sU.jpg?s=328&g=1",
		//"-w",
		//"../../testdata/mergi_logo_watermark.png 0 0",
		"-i",
		"https://i.stack.imgur.com/Ni7sU.jpg?s=328&g=1",
		//"-w",
		//"../../testdata/mergi_logo_watermark.png 0 0",
		"-i",
		"https://i.stack.imgur.com/Ni7sU.jpg?s=328&g=1",
		//"-w",
		//"../../testdata/mergi_logo_watermark.png 0 0",
		"-i",
		"https://i.stack.imgur.com/Ni7sU.jpg?s=328&g=1",
		"-w",
		"../../testdata/mergi_logo_watermark.png 0 0",
		//"-w",
		//"../../testdata/mergi_logo_watermark.png 0 0",
		"-f",
		"true",
		//"-c",
		//"0 0 300 328",
		"-f",
		"true",
		"-w",
		"../../testdata/mergi_logo_watermark.png -r 180 80 0 0",
	}
	os.Args = args
	main()
}

func TestSpriteFoodAnimation(t *testing.T) {
	args2 := []string{
		"mergi",
		"-t",
		"TBTBTB",

		"-i", "../../testdata/avocado-3210885_960_720.jpg",
		"-r", "480 360",

		"-i", "../../testdata/cherry-3074284_960_720.jpg",
		"-r", "480 360",

		"-i", "../../testdata/grapes-2032838_960_720.jpg",
		"-r", "480 360",

		"-i", "../../testdata/pomegranate-3259161_960_720.jpg",
		"-r", "480 360",

		"-i", "../../testdata/tangerine-2914928_960_720.jpg",
		"-r", "480 360",

		"-i", "../../testdata/smoothie-3193660_960_720.jpg",
		"-r", "480 360",

		"-a", "sprite 10",

		"-o", "out.gif",
	}
	os.Args = args2
	main()
}

func Test2x2FoodsScaleMain(t *testing.T) {
	args2 := []string{
		"mergi",
		"-t",
		"TBTBTB",

		"-i", "../../testdata/avocado-3210885_960_720.jpg",
		"-r", "480 360",

		"-i", "../../testdata/cherry-3074284_960_720.jpg",
		"-r", "480 360",

		"-i", "../../testdata/grapes-2032838_960_720.jpg",
		"-r", "480 360",

		"-i", "../../testdata/pomegranate-3259161_960_720.jpg",
		"-r", "480 360",

		"-i", "../../testdata/tangerine-2914928_960_720.jpg",
		"-r", "480 360",

		"-i", "../../testdata/smoothie-3193660_960_720.jpg",
		"-r", "480 360",

		"-f", "true",

		"-r", "2880 1440",

		"-w", "../../testdata/mergi_logo_watermark.png -r 180 80 2300 150",

		"-o", "out.png",
	}
	os.Args = args2
	main()
}

func Test2x2FruitsMain(t *testing.T) {
	args2 := []string{
		"mergi",
		"-t", "TBTB",
		"-i", "https://cdn.pixabay.com/photo/2017/02/02/14/04/grapes-2032838_960_720.jpg",
		"-i", "https://cdn.pixabay.com/photo/2018/01/10/16/59/cherry-3074284_960_720.jpg",
		"-i", "https://cdn.pixabay.com/photo/2014/06/18/13/25/cherries-371233_960_720.jpg",
		"-i", "https://cdn.pixabay.com/photo/2016/03/05/23/01/berry-1239429_960_720.jpg",
		"-f", "true",

		"-w", "../../testdata/mergi_logo_watermark.png 1550 1000",

		"-o", "out.png",
	}
	os.Args = args2
	main()
}

func Test2x2AnimalsMain(t *testing.T) {
	args2 := []string{
		"mergi",
		"-t", "TBTB",

		"-i", "../../testdata/hedgehog-child-3636026_960_720.jpg",
		"-c", "400 100 300 300",

		"-i", "../../testdata/lion-3576045_960_720.jpg",
		"-c", "400 100 300 300",

		"-i", "../../testdata/rabbit-1882699_960_720.jpg",
		"-c", "400 100 300 300",

		"-i", "../../testdata/tiger-2320819_960_720.jpg",
		"-c", "200 100 300 300",

		"-f", "true",

		"-w", "../../testdata/mergi_logo_watermark.png -r 90 40 500 560",

		"-o", "out.png",
	}
	os.Args = args2
	main()
}

func TestCropMain(t *testing.T) {
	args2 := []string{
		"mergi",
		"-i", "https://cdn.pixabay.com/photo/2015/06/08/15/02/pug-801826_960_720.jpg",
		"-c", "410 50 130 200",
		"-o", "out.png",
	}
	os.Args = args2
	main()
}

func TestCropFoodMain(t *testing.T) {
	args2 := []string{
		"mergi",
		"-t", "TTT",
		"-i", "https://cdn.pixabay.com/photo/2014/06/11/17/00/cook-366875__340.jpg",
		"-c", "450 200 110 130",
		"-i", "https://cdn.pixabay.com/photo/2014/06/11/17/00/cook-366875__340.jpg",
		"-c", "340 200 110 130",
		"-i", "https://cdn.pixabay.com/photo/2014/06/11/17/00/cook-366875__340.jpg",
		"-c", "230 200 110 130",
	}
	os.Args = args2
	main()
}

func TestCropFacesMain(t *testing.T) {
	args2 := []string{
		"mergi",
		"-t", "TBTB",
		"-i", "https://cdn.pixabay.com/photo/2016/03/09/09/22/workplace-1245776_960_720.jpg",
		"-c", "450 220 110 130",
		"-i", "https://cdn.pixabay.com/photo/2016/03/09/09/22/workplace-1245776_960_720.jpg",
		"-c", "80 210 110 130",
		"-i", "https://cdn.pixabay.com/photo/2016/03/09/09/22/workplace-1245776_960_720.jpg",
		"-c", "110 100 110 130",
		"-i", "https://cdn.pixabay.com/photo/2016/03/09/09/22/workplace-1245776_960_720.jpg",
		"-c", "540 100 110 130",
	}
	os.Args = args2
	main()
}

func TestResizeDogMain(t *testing.T) {
	args2 := []string{
		"mergi",
		"-i", "https://cdn.pixabay.com/photo/2015/03/26/09/54/pug-690566_960_720.jpg",
		"-r", "480 320",
	}
	os.Args = args2
	main()
}

func TestCropResizeMain(t *testing.T) {
	args2 := []string{
		"mergi",
		"-i", "https://cdn.pixabay.com/photo/2016/03/26/13/09/cup-of-coffee-1280537_960_720.jpg",
		"-c", "530 70 380 350",
		"-r", "760 700",
	}
	os.Args = args2
	main()
}

func TestCropResizeMergeMain(t *testing.T) {
	args2 := []string{
		"mergi",
		"-t", "TTBB",
		"-i", "https://cdn.pixabay.com/photo/2018/03/31/06/31/dog-3277416_960_720.jpg",
		"-c", "300 0 400 640",
		"-i", "https://cdn.pixabay.com/photo/2016/07/15/15/55/dachshund-1519374_960_720.jpg",
		"-c", "270 0 550 550",
		"-r", "225 225",
		"-i", "https://cdn.pixabay.com/photo/2016/02/26/16/32/dog-1224267_960_720.jpg",
		"-c", "220 50 550 550",
		"-r", "225 225",
		"-i", "https://cdn.pixabay.com/photo/2016/05/09/10/42/weimaraner-1381186_960_720.jpg",
		"-c", "220 50 550 550",
		"-r", "225 225",
		"-f", "true",
		"-r", "1280 1280",
		"-w", "../../testdata/mergi_logo_watermark.png 0 0",
	}
	os.Args = args2
	main()
}

func TestFloatLeftAnimalsMain2(t *testing.T) {
	args2 := []string{
		"mergi",
		"-t", "TTTT",

		"-i", "https://cdn.pixabay.com/photo/2018/08/27/21/36/hedgehog-child-3636026_960_720.jpg",
		"-c", "400 0 300 600",

		"-i", "https://cdn.pixabay.com/photo/2018/07/31/22/08/lion-3576045_960_720.jpg",
		"-c", "400 0 300 600",

		"-i", "https://cdn.pixabay.com/photo/2016/12/04/21/58/rabbit-1882699_960_720.jpg",
		"-c", "400 0 300 600",

		"-i", "https://cdn.pixabay.com/photo/2017/05/17/12/42/tiger-2320819_960_720.jpg",
		"-c", "200 0 300 600",

		"-f", "true",

		"-w", "../../testdata/mergi_logo_watermark.png -r 180 80 1020 0",

		"-o", "out.png",
	}
	os.Args = args2
	main()
}

func TestFloatLeftAnimalsMain3(t *testing.T) {
	args2 := []string{
		"mergi",
		"-t", "TTBTBT",

		"-i", "https://cdn.pixabay.com/photo/2016/01/03/17/59/bananas-1119790_960_720.jpg",
		"-c", "400 0 300 600",

		"-i", "https://cdn.pixabay.com/photo/2015/03/26/09/40/blueberries-690072_960_720.jpg",
		"-c", "400 0 300 300",

		"-i", "https://cdn.pixabay.com/photo/2018/05/08/20/19/pomegranate-3383814_960_720.jpg",
		"-c", "400 0 300 300",

		"-i", "https://cdn.pixabay.com/photo/2012/02/19/18/05/oranges-15046_960_720.jpg",
		"-c", "400 300 300 300",

		"-i", "https://cdn.pixabay.com/photo/2016/08/10/15/27/green-1583575_960_720.jpg",
		"-c", "100 200 300 300",

		"-i", "https://cdn.pixabay.com/photo/2013/04/02/21/47/strawberries-99551_960_720.jpg",
		"-c", "200 0 300 600",

		"-f", "true",

		"-w", "../../testdata/mergi_logo_watermark.png -r 180 80 1020 0",

		"-o", "out.png",
	}
	os.Args = args2
	main()
}

func TestFloatLeftAnimalsMain(t *testing.T) {
	args2 := []string{
		"mergi",
		"-t", "TTTT",

		"-i", "../../testdata/hedgehog-child-3636026_960_720.jpg",
		"-c", "400 0 300 600",

		"-i", "../../testdata/lion-3576045_960_720.jpg",
		"-c", "400 0 300 600",

		"-i", "../../testdata/rabbit-1882699_960_720.jpg",
		"-c", "400 0 300 600",

		"-i", "../../testdata/tiger-2320819_960_720.jpg",
		"-c", "200 0 300 600",

		"-f", "true",

		"-w", "https://i.imgur.com/S4MRB3d.png -r 180 80 900 0",

		"-o", "out.png",
	}
	os.Args = args2
	main()
}
