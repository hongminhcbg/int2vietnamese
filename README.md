# int2vietnamese
Convert int to VietNamese string for golang

Result 
    
    250331944655756829 	=> 	hai trăm năm mươi tỷ tỷ ba trăm ba mươi mốt nghìn tỷ chín trăm bốn mươi tư tỷ sáu trăm năm mươi lăm triệu bảy trăm năm mươi sáu nghìn tám trăm hai mươi chín

    280612196147269722 	=> 	hai trăm tám mươi tỷ tỷ sáu trăm mười hai nghìn tỷ một trăm chín mươi sáu tỷ một trăm bốn mươi bảy triệu hai trăm sáu mươi chín nghìn bảy trăm hai mươi hai

    191184799421266057 	=> 	một trăm chín mươi mốt tỷ tỷ một trăm tám mươi tư nghìn tỷ bảy trăm chín mươi chín tỷ bốn trăm hai mươi mốt triệu hai trăm sáu mươi sáu nghìn không trăm năm mươi bảy

    798601829110992793 	=> 	bảy trăm chín mươi tám tỷ tỷ sáu trăm lẻ một nghìn tỷ tám trăm hai mươi chín tỷ một trăm mười triệu chín trăm chín mươi hai nghìn bảy trăm chín mươi ba

    925808077690296198 	=> 	chín trăm hai mươi lăm tỷ tỷ tám trăm lẻ tám nghìn tỷ không trăm bảy mươi bảy tỷ sáu trăm chín mươi triệu hai trăm chín mươi sáu nghìn một trăm chín mươi tám

    506669535813863363 	=> 	năm trăm lẻ sáu tỷ tỷ sáu trăm sáu mươi chín nghìn tỷ năm trăm ba mươi lăm tỷ tám trăm mười ba triệu tám trăm sáu mươi ba nghìn ba trăm sáu mươi ba

    117127360716071582 	=> 	một trăm mười bảy tỷ tỷ một trăm hai mươi bảy nghìn tỷ ba trăm sáu mươi tỷ bảy trăm mười sáu triệu không trăm bảy mươi mốt nghìn năm trăm tám mươi hai
    
    81501739640992003 	=> 	tám mươi mốt tỷ tỷ năm trăm lẻ một nghìn tỷ bảy trăm ba mươi chín tỷ sáu trăm bốn mươi triệu chín trăm chín mươi hai nghìn không trăm lẻ ba

    610519428235075200 	=> 	sáu trăm mười tỷ tỷ năm trăm mười chín nghìn tỷ bốn trăm hai mươi tám tỷ hai trăm ba mươi lăm triệu không trăm bảy mươi lăm nghìn hai trăm

    832421759198250571 	=> 	tám trăm ba mươi hai tỷ tỷ bốn trăm hai mươi mốt nghìn tỷ bảy trăm năm mươi chín tỷ một trăm chín mươi tám triệu hai trăm năm mươi nghìn năm trăm bảy mươi mốt

    650787758596762 	=> 	sáu trăm năm mươi nghìn tỷ bảy trăm tám mươi bảy tỷ bảy trăm năm mươi tám triệu năm trăm chín mươi sáu nghìn bảy trăm sáu mươi hai

    309781370579293 	=> 	ba trăm lẻ chín nghìn tỷ bảy trăm tám mươi mốt tỷ ba trăm bảy mươi triệu năm trăm bảy mươi chín nghìn hai trăm chín mươi ba

    540288371717485 	=> 	năm trăm bốn mươi nghìn tỷ hai trăm tám mươi tám tỷ ba trăm bảy mươi mốt triệu bảy trăm mười bảy nghìn bốn trăm tám mươi lăm

    357378392750796 	=> 	ba trăm năm mươi bảy nghìn tỷ ba trăm bảy mươi tám tỷ ba trăm chín mươi hai triệu bảy trăm năm mươi nghìn bảy trăm chín mươi sáu

    162444015041355 	=> 	một trăm sáu mươi hai nghìn tỷ bốn trăm bốn mươi tư tỷ không trăm mười lăm triệu không trăm bốn mươi mốt nghìn ba trăm năm mươi lăm

    77170949307269 	=> 	bảy mươi bảy nghìn tỷ một trăm bảy mươi tỷ chín trăm bốn mươi chín triệu ba trăm lẻ bảy nghìn hai trăm sáu mươi chín

    360900136436604 	=> 	ba trăm sáu mươi nghìn tỷ chín trăm tỷ một trăm ba mươi sáu triệu bốn trăm ba mươi sáu nghìn sáu trăm lẻ bốn

    996583140974077 	=> 	chín trăm chín mươi sáu nghìn tỷ năm trăm tám mươi ba tỷ một trăm bốn mươi triệu chín trăm bảy mươi tư nghìn không trăm bảy mươi bảy

    403262568062214 	=> 	bốn trăm lẻ ba nghìn tỷ hai trăm sáu mươi hai tỷ năm trăm sáu mươi tám triệu không trăm sáu mươi hai nghìn hai trăm mười bốn

    578932389288766 	=> 	năm trăm bảy mươi tám nghìn tỷ chín trăm ba mươi hai tỷ ba trăm tám mươi chín triệu hai trăm tám mươi tám nghìn bảy trăm sáu mươi sáu

    411976117745 	=> 	bốn trăm mười một tỷ chín trăm bảy mươi sáu triệu một trăm mười bảy nghìn bảy trăm bốn mươi lăm

    665431694775 	=> 	sáu trăm sáu mươi lăm tỷ bốn trăm ba mươi mốt triệu sáu trăm chín mươi tư nghìn bảy trăm bảy mươi lăm

    401735430368 	=> 	bốn trăm lẻ một tỷ bảy trăm ba mươi lăm triệu bốn trăm ba mươi nghìn ba trăm sáu mươi tám

    590270084178 	=> 	năm trăm chín mươi tỷ hai trăm bảy mươi triệu không trăm tám mươi tư nghìn một trăm bảy mươi tám

    25357429797 	=> 	hai mươi lăm tỷ ba trăm năm mươi bảy triệu bốn trăm hai mươi chín nghìn bảy trăm chín mươi bảy

    689600554401 	=> 	sáu trăm tám mươi chín tỷ sáu trăm triệu năm trăm năm mươi tư nghìn bốn trăm lẻ một

    742602364957 	=> 	bảy trăm bốn mươi hai tỷ sáu trăm lẻ hai triệu ba trăm sáu mươi tư nghìn chín trăm năm mươi bảy

    939138918603 	=> 	chín trăm ba mươi chín tỷ một trăm ba mươi tám triệu chín trăm mười tám nghìn sáu trăm lẻ ba

    685166188614 	=> 	sáu trăm tám mươi lăm tỷ một trăm sáu mươi sáu triệu một trăm tám mươi tám nghìn sáu trăm mười 
    bốn

    160601519670 	=> 	một trăm sáu mươi tỷ sáu trăm lẻ một triệu năm trăm mười chín nghìn sáu trăm bảy mươi

    393035473 		=> 	ba trăm chín mươi ba triệu không trăm ba mươi lăm nghìn bốn trăm bảy mươi ba

    5152351 		=> 	năm triệu một trăm năm mươi hai nghìn ba trăm năm mươi mốt

    804770696 		=> 	tám trăm lẻ bốn triệu bảy trăm bảy mươi nghìn sáu trăm chín mươi sáu

    283976832 		=> 	hai trăm tám mươi ba triệu chín trăm bảy mươi sáu nghìn tám trăm ba mươi hai

    24406493 		=> 	hai mươi tư triệu bốn trăm lẻ sáu nghìn bốn trăm chín mươi ba

    679937663 		=> 	sáu trăm bảy mươi chín triệu chín trăm ba mươi bảy nghìn sáu trăm sáu mươi ba

    1647702 		=> 	một triệu sáu trăm bốn mươi bảy nghìn bảy trăm lẻ hai

    96366253 		=> 	chín mươi sáu triệu ba trăm sáu mươi sáu nghìn hai trăm năm mươi ba

    189075545 		=> 	một trăm tám mươi chín triệu không trăm bảy mươi lăm nghìn năm trăm bốn mươi lăm

    696951521 		=> 	sáu trăm chín mươi sáu triệu chín trăm năm mươi mốt nghìn năm trăm hai mươi mốt

    458955 		=> 	bốn trăm năm mươi tám nghìn chín trăm năm mươi lăm

    11034 		=> 	mười một nghìn không trăm ba mươi tư

    709406 		=> 	bảy trăm lẻ chín nghìn bốn trăm lẻ sáu

    395340 		=> 	ba trăm chín mươi lăm nghìn ba trăm bốn mươi

    762358 		=> 	bảy trăm sáu mươi hai nghìn ba trăm năm mươi tám

    364707 		=> 	ba trăm sáu mươi tư nghìn bảy trăm lẻ bảy

    680399 		=> 	sáu trăm tám mươi nghìn ba trăm chín mươi chín

    270109 		=> 	hai trăm bảy mươi nghìn một trăm lẻ chín

    623452 		=> 	sáu trăm hai mươi ba nghìn bốn trăm năm mươi hai

    245750 		=> 	hai trăm bốn mươi lăm nghìn bảy trăm năm mươi

    126 		=> 	một trăm hai mươi sáu

    373 		=> 	ba trăm bảy mươi ba

    301 		=> 	ba trăm lẻ một

    243 		=> 	hai trăm bốn mươi ba

    758 		=> 	bảy trăm năm mươi tám

    150 		=> 	một trăm năm mươi

    234 		=> 	hai trăm ba mươi tư

    369 		=> 	ba trăm sáu mươi chín

    860 		=> 	tám trăm sáu mươi
    
    17 		    => 	mười bảy



Example:

    package main

    import (
    	"fmt"
    
    	int2vietnamese "github.com/hongminhcbg/int2vietnamese"
    )
    
    func main() {
    	fmt.Println(int2vietnamese.Int2VietNamese(100))
    }

      
