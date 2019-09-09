package hash_table

import (
	"testing"
	"fmt"
)

func TestLongestPalindrome(t *testing.T) {
	s := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
	v := longestPalindrome(s)
	fmt.Println(v)
}

func TestFindWords(t *testing.T)  {
	s := []string{"Hello", "Alaska", "Dad", "Peace"}
	v := findWords(s)
	fmt.Println(v)
}

func TestFindLHS(t *testing.T) {
	s := []int{1,1,1,1,1}
	v := findLHS(s)
	fmt.Println(v)
}

func TestMyHashMap(t *testing.T) {
	mk := map[int]int{}
	mv := map[int]int{}
	obj := ConstructorMap()
	action := []string{"put","put","put","put","remove","put","put","put","put","put","put","get","put","get","put","remove","get","put","put","put","put","get","remove","get","remove","get","put","put","put","put","put","put","remove","remove","put","put","remove","get","put","put","put","remove","put","get","put","get","put","remove","put","put","get","remove","put","put","remove","put","put","put","put","put","remove","put","remove","put","put","put","put","remove","put","get","remove","put","put","put","put","put","put","put","put","remove","remove","put","put","put","get","put","put","put","remove","remove","put","put","remove","remove","put","put","put","remove","remove","put","put","put","put","put","put","remove","get","put","get","get","put","put","remove","put","get","put","get","put","put","put","put","put","remove","put","remove","put","put","put","get","put","put","put","put","put","put","put","get","put","put","put","put","put","remove","remove","remove","put","put","put","put","put","get","put","put","put","get","get","remove","put","put","put","put","remove","get","get","put","put","get","remove","get","put","put","put","remove","put","get","put","get","put","put","put","remove","put","remove","put","put","put","put","remove","put","put","put","put","remove","put","put","put","put","put","put","put","remove","remove","remove","put","put","put","get","put","put","put","put","remove","get","put","get","put","put","remove","put","put","put","put","put","put","put","remove","remove","remove","put","remove","get","put","get","put","get","put","put","put","put","put","get","remove","put","put","put","put","remove","put","put","get","put","put","put","remove","put","put","put","put","remove","put","remove","remove","put","put","put","remove","put","put","put","put","put","get","put","get","put","get","get","put","put","remove","put","remove","put","put","put","put","remove","get","get","remove","put","put","put","remove","get","get","put","put","put","put","put","remove","get","put","put","put","get","put","put","remove","put","put","put","get","get","put","put","put","remove","put","get","put","put","remove","put","remove","remove","put","put","put","put","get","put","put","put","remove","remove","put","get","put","put","put","put","put","remove","put","put","put","get","remove","put","remove","put","put","put","remove","put","put","put","remove","put","remove","put","get","put","put","remove","put","put","put","put","put","put","put","put","put","put","put","put","put","get","put","get","get","put","get","remove","remove","put","put","put","put","put","put","put","remove","get","put","put","put","put","put","put","put","remove","put","put","put","get","get","get","put","get","put","put","get","put","put","put","remove","put","put","put","put","put","remove","put","put","put","put","get","put","put","put","put","remove","put","put","remove","put","put","remove","get","put","put","get","get","get","put","put","put","get","remove","put","get","remove","put","get","put","get","get","put","remove","put","put","put","put","remove","put","put","put","put","get","put","put","get","get","put","put","get","get","put","put","put","get","put","put","remove","remove","put","put","put","remove","remove","put","put","put","put","put","put","put","put","put","put","put","put","get","put","put","put","get","put","put","remove","put","put","put","put","get","get","put","put","remove","put","put","put","put","put","get","put","get","get","put","put","put","put","put","put","get","put","remove","put","put","put","put","put","put","put","put","put","put","put","remove","put","get","put","put","get","put","get","put","put","put","put","get","remove","put","remove","get","remove","put","put","put","put","put","get","put","put","put","put","get","put","put","put","remove","put","remove","put","put","put","put","get","remove","remove","put","put","put","put","put","put","put","put","put","remove","put","put","put","put","get","put","put","remove","put","remove","get","put","put","put","get","get","remove","put","put","remove","put","put","put","put","put","get","put","put","remove","get","get","put","remove","put","put","get","get","put","get","put","remove","put","put","put","remove","put","put","put","put","put","put","put","put","put","put","put","put","remove","put","get","put","put","put","get","get","remove","put","put","remove","get","put","get","remove","put","put","put","put","put","put","put","put","put","remove","put","put","remove","get","get","remove","put","remove","put","put","get","put","put","put","put","get","put","get","put","remove","put","put","put","put","put","put","get","get","put","put","put","put","put","get","get","put","remove","remove","put","remove","get","get","put","put","put","remove","remove","get","get","remove","put","put","put","get","put","put","put","put","put","put","get","remove","put","put","put","put","get","put","remove","put","get","get","get","put","put","put","put","put","get","get","put","put","put","put","put","put","put","put","put","put","put","get","put","put","get","get","put","get","put","put","put","put","put","put","get","put","get","put","put","put","remove","get","get","put","put","remove","put","put","put","get","put","put","get","put","put","get","put","put","put","remove","put","put","remove","put","put","remove","put","put","put","put","get","get","get","put","put","put","put","get","remove","put","put","remove","get","get","put","put","get","put","get","get","put","put","put","put","get","remove","put","put","put","put","get","put","put","remove","put","put","put","put","put","put","put","get","put","remove","put","put","put","put","get","put","put","put","get","put","put","put","get","put","put","put","put","put","get","put","put","remove","put","put","get","put","put","put","put","get","put","put","get","put","get","put","put","put","put","remove","put","put","put","put","put","remove","remove","put","put","remove","put","put","put","put","put","get","remove","get","put","get","remove","put","get","remove","get","put","put","put","put","put","put","put","put","remove","remove","get","put","put","remove","put","put","put","put","put","put","put","get","put","put","put","put","put","put","put","remove","get","remove","remove","put","put","get","put","put","put","put","put","remove","put","put","remove","put","put","put","put","put","put","remove","remove","remove","put","put","put","get","put","put","put","remove","put","put"}

	params := [][]int{{970,538},{908,29},{395,865},{250,847},{836},{233,568},{657,790},{595,271},{769,219},{55,112},{157,493},{920},{632,358},{669},{506,228},{904},{473},{461,40},{748,973},{446,544},{766,461},{395},{211},{415},{157},{252},{252,22},{942,681},{600,988},{424,39},{685,482},{561,605},{632},{461},{916,329},{739,735},{769},{942},{460,226},{183,411},{224,524},{769},{508,614},{632},{254,825},{603},{115,667},{460},{940,813},{50,629},{519},{595},{39,913},{742,13},{765},{163,627},{554,130},{255,945},{22,78},{912,390},{632},{609,410},{956},{515,243},{975,871},{520,313},{682,538},{563},{119,902},{916},{766},{293,885},{657,665},{518,832},{897,489},{340,972},{790,24},{637,445},{544,498},{115},{600},{269,209},{734,3},{243,108},{233},{679,632},{640,55},{288,301},{682},{871},{922,755},{624,787},{776},{293},{564,902},{32,743},{934,278},{250},{389},{620,711},{420,623},{346,959},{829,832},{776,894},{465,769},{508},{358},{126,481},{255},{50},{477,991},{973,337},{32},{823,283},{21},{733,431},{583},{735,407},{873,702},{578,256},{813,221},{669,432},{790},{941,945},{645},{560,775},{823,772},{458,220},{243},{947,136},{168,560},{946,19},{172,608},{624,260},{876,516},{76,334},{704},{615,737},{110,453},{189,678},{746,201},{497,330},{632},{993},{497},{915,545},{329,558},{662,773},{208,135},{797,614},{640},{108,809},{262,401},{560,965},{494},{222},{922},{237,439},{688,796},{974,331},{367,470},{339},{115},{790},{658,873},{513,520},{2},{76},{110},{199,511},{674,853},{546,849},{430},{394,41},{189},{443,983},{270},{353,44},{959,995},{275,798},{910},{522,939},{176},{520,292},{631,538},{216,551},{197,623},{55},{452,416},{805,462},{884,223},{484,454},{126},{59,954},{454,575},{500,940},{999,520},{211,106},{190,915},{995,462},{519},{368},{734},{549,845},{47,885},{609,356},{765},{163,520},{382,774},{190,545},{146,166},{748},{282},{357,531},{113},{89,489},{994,5},{967},{972,826},{220,91},{296,808},{826,7},{401,910},{126,802},{534,781},{293},{282},{515},{29,473},{390},{32},{265,801},{658},{29,648},{678},{40,493},{9,809},{701,244},{345,542},{386,296},{632},{89},{575,758},{168,575},{286,619},{210,335},{546},{77,534},{380,885},{414},{590,272},{924,933},{662,1},{544},{349,402},{244,274},{907,303},{638,453},{339},{255,802},{973},{130},{266,695},{486,444},{557,522},{518},{392,750},{538,174},{407,337},{455,912},{755,278},{242},{947,647},{367},{657,280},{286},{255},{307,453},{382,216},{609},{69,582},{340},{36,792},{4,288},{518,878},{376,516},{873},{211},{417},{908},{695,90},{646,513},{114,675},{924},{321},{486},{300,87},{68,223},{924,665},{680,323},{410,97},{538},{685},{884,617},{496,967},{155,853},{367},{694,765},{607,567},{748},{565,855},{151,633},{300,838},{76},{163},{796,591},{516,514},{95,921},{353},{445,40},{32},{163,846},{697,892},{637},{870,672},{560},{410},{909,414},{319,973},{259,144},{423,142},{165},{984,652},{121,133},{117,13},{340},{390},{237,305},{520},{712,752},{380,7},{661,937},{249,9},{937,291},{168},{501,89},{989,540},{736,722},{739},{250},{741,270},{920},{421,848},{917,716},{788,294},{520},{642,352},{250,775},{710,268},{349},{801,427},{921},{905,183},{69},{163,429},{274,863},{108},{410,289},{578,258},{613,113},{992,527},{155,194},{557,696},{190,519},{630,59},{234,795},{740,324},{270,492},{107,63},{705,446},{205},{390,676},{374},{557},{321,754},{710},{121},{488},{156,679},{132,238},{793,469},{966,912},{612,691},{860,749},{245,736},{829},{237},{858,449},{655,591},{913,760},{977,737},{13,954},{189,252},{215,933},{168},{682,867},{802,747},{710,342},{994},{257},{288},{972,139},{255},{637,821},{383,858},{685},{146,53},{464,179},{98,642},{911},{524,259},{196,641},{644,843},{271,723},{706,285},{205},{256,884},{291,722},{750,142},{205,50},{259},{816,839},{123,741},{913,349},{894,577},{914},{129,267},{242,60},{288},{189,30},{507,805},{277},{227},{332,850},{716,309},{493},{401},{669},{682,330},{880,925},{4,287},{858},{177},{351,217},{621},{193},{611,128},{183},{789,817},{655},{658},{501,965},{412},{334,651},{345,411},{73,9},{162,791},{644},{330,608},{998,312},{96,234},{906,419},{221},{179,324},{507,978},{632},{134},{155,691},{92,442},{353},{861},{714,403},{824,261},{814,164},{118},{931,816},{39,844},{790},{894},{980,797},{673,80},{109,321},{293},{256},{625,277},{272,650},{890,666},{282,399},{50,95},{635,169},{25,749},{681,759},{126,84},{229,309},{892,941},{505,343},{496},{843,820},{259,362},{549,86},{340},{602,374},{711,73},{196},{519,269},{785,465},{747,706},{449,878},{407},{357},{459,357},{615,903},{191},{86,672},{467,638},{406,351},{665,29},{416,514},{739},{563,810},{190},{402},{705,575},{91,324},{993,183},{902,528},{136,560},{866,309},{458},{13,751},{999},{201,651},{666,246},{969,696},{249,255},{632,413},{109,432},{866,745},{203,476},{72,869},{379,178},{955,758},{367},{864,228},{13},{101,120},{679,497},{814},{216,235},{401},{588,717},{914,836},{648,29},{253,926},{316},{69},{978,205},{963},{701},{941},{758,489},{506,806},{597,467},{618,861},{863,454},{190},{490,703},{866,151},{735,303},{217,552},{460},{850,691},{554,450},{1,997},{424},{603,68},{319},{883,877},{489,109},{819,400},{957,400},{467},{601},{481},{832,895},{979,52},{137,90},{394,540},{198,303},{562,223},{680,851},{911,399},{730,427},{563},{983,9},{559,525},{422,86},{121,720},{531},{151,902},{468,501},{39},{932,29},{650},{598},{131,715},{752,133},{119,831},{789},{606},{941},{735,652},{107,448},{98},{436,737},{728,947},{533,835},{971,356},{776,984},{880},{939,191},{119,844},{168},{733},{39},{739,760},{940},{265,768},{523,176},{969},{269},{464,948},{456},{183,358},{559},{211,986},{696,242},{427,90},{785},{868,968},{976,307},{667,619},{798,831},{297,468},{802,53},{221,956},{319,908},{791,313},{56,158},{494,398},{387,909},{298},{180,573},{839},{663,316},{460,628},{674,243},{254},{922},{197},{814,245},{560,195},{392},{788},{439,556},{698},{746},{138,331},{379,700},{984,296},{937,471},{968,888},{220,87},{706,368},{430,702},{466,349},{441},{334,642},{324,670},{635},{406},{734},{611},{954,130},{666},{601,106},{951,372},{382},{64,691},{481,581},{91,807},{113,529},{307},{129,955},{383},{926,412},{644},{86,862},{879,202},{661,959},{849,806},{583,501},{331,825},{922},{391},{68,504},{149,731},{463,64},{572,623},{165,760},{268},{448},{819,282},{586},{45},{353,316},{589},{655},{249},{382,367},{963,184},{625,538},{932},{996},{935},{612},{788},{252,397},{728,935},{765,924},{976},{784,464},{806,326},{625,866},{512,83},{569,768},{473,714},{401},{683},{603,271},{762,849},{449,150},{242,372},{285},{31,474},{716},{653,344},{673},{588},{813},{672,111},{120,299},{353,497},{297,510},{619,677},{390},{50},{185,215},{809,322},{749,218},{868,920},{955,207},{530,64},{95,906},{940,961},{505,893},{339,428},{364,122},{426},{646,775},{455,542},{423},{807},{142,73},{392},{174,439},{575,427},{321,697},{9,998},{32,272},{351,496},{834},{481,71},{760},{737,574},{812,451},{626,196},{102},{648},{617},{166,152},{701,720},{750},{625,425},{335,404},{820,328},{240},{226,815},{508,265},{785},{863,881},{298,263},{325},{12,960},{439,28},{502,331},{114},{117,930},{922,65},{456},{686,211},{668,478},{601},{839,336},{737,282},{756,191},{581,547},{89},{167},{695},{272,244},{449,570},{688,413},{142,769},{382},{522},{794,632},{714,983},{163},{737},{310},{771,745},{254,395},{551},{765,622},{263},{751},{986,924},{76,585},{548,618},{589,427},{785},{866},{541,789},{58,900},{816,495},{324,395},{626},{648,807},{280,557},{92},{994,295},{852,148},{700,192},{626,551},{628,937},{442,892},{268,263},{536},{593,528},{567},{164,733},{997,800},{583,941},{967,894},{737},{825,103},{528,774},{485,373},{611},{935,118},{190,262},{494,836},{917},{997,448},{269,659},{941,894},{913,554},{63,246},{167},{235,456},{136,501},{309},{212,503},{796,931},{337},{129,759},{302,276},{360,779},{809,918},{55},{888,433},{535,160},{97},{769,627},{762},{449,596},{703,27},{133,407},{509,67},{266},{697,778},{903,748},{430,909},{464,626},{452,493},{891},{293},{20,945},{686,88},{329},{207,827},{202,641},{709,186},{303,451},{385,541},{409},{201},{766},{395,956},{971},{704},{250,34},{354},{766},{789},{106,974},{652,301},{654,875},{109,363},{481,935},{361,686},{989,931},{488,673},{979},{806},{450},{509,414},{528,835},{673},{328,121},{303,406},{45,174},{410,906},{42,356},{675,953},{882,222},{978},{252,396},{347,94},{367,754},{578,278},{107,628},{73,583},{166,100},{162},{593},{791},{740},{697,977},{368,152},{971},{918,197},{152,444},{995,814},{869,900},{335,818},{250},{951,948},{605,293},{959},{873,243},{236,239},{577,636},{244,441},{884,904},{942,993},{708},{581},{82},{293,859},{622,486},{878,920},{220},{493,182},{445,582},{392,227},{523},{800,270},{404,803}}

	result := []int{-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,865,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,681,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,329,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,568,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,945,629,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,108,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,55,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,453,-1,-1,-1,-1,-1,678,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,873,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,470,-1,619,802,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,106,-1,-1,-1,-1,-1,-1,-1,444,-1,-1,-1,-1,-1,-1,482,-1,-1,-1,470,-1,-1,-1,-1,-1,-1,-1,520,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,292,-1,-1,-1,-1,-1,-1,-1,-1,-1,735,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,582,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,696,-1,268,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,305,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,5,-1,301,-1,802,-1,-1,482,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,144,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,910,432,-1,-1,-1,449,-1,-1,-1,-1,-1,411,-1,591,873,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,967,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,337,531,-1,-1,-1,-1,-1,-1,-1,-1,735,-1,519,-1,-1,-1,-1,-1,-1,-1,220,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,751,-1,-1,164,-1,910,-1,-1,-1,-1,-1,-1,-1,-1,244,-1,-1,-1,-1,-1,-1,519,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,638,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,817,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,925,-1,-1,-1,431,-1,-1,-1,-1,-1,696,209,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,825,-1,-1,-1,-1,-1,294,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,351,-1,-1,-1,-1,-1,-1,216,-1,-1,-1,-1,453,-1,858,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,591,255,-1,-1,-1,-1,-1,-1,691,-1,-1,-1,-1,307,-1,-1,-1,-1,-1,-1,910,-1,-1,-1,-1,-1,-1,-1,-1,-1,80,717,221,-1,-1,-1,-1,-1,676,95,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,142,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,29,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,90,-1,-1,-1,-1,367,-1,-1,-1,-1,282,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,196,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,282,-1,-1,-1,-1,-1,-1,-1,716,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,849,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,356,-1,-1,-1,-1,817,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,205,-1,-1,-1,-1,-1,-1,-1,-1,528,-1,-1,-1,-1,356,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,87,-1,-1,-1,-1,-1,-1}

	fmt.Println(len(action))
	fmt.Println(len(params))
	fmt.Println(len(result))

	for i, v := range action {

		switch v {
		case "put":
			obj.Put(params[i][0], params[i][1])
			mk[params[i][0]] ++
			mv[params[i][1]] ++
			break
		case "remove":
			obj.Remove(params[i][0])
			break
		case "get":
			v:= obj.Get(params[i][0])
			fmt.Println("i:%i, res:%b, p, res, v, node", i,v == result[i], params[i], result[i], v)
			break
		}
	}

	fmt.Println(obj.list)
	fmt.Println(mk)
	fmt.Println(mv)
}


func TestMyHashMapOne(t *testing.T) {
	obj := ConstructorMap()
	obj.Put(519, 668)
	obj.Put(115, 667)
	obj.Put(115, 669)
	v:= obj.Get(115)
	v1 := obj.Get(519)
	fmt.Println(v)
	fmt.Println(v1)
}

func TestLongestWord (t *testing.T) {
	 // ["w","wo","wor","worl","world"]
	 v := longestWord([]string{"ts","e","x","pbhj","opto","xhigy","erikz","pbh","opt","erikzb","eri","erik","xlye","xhig","optoj","optoje","xly","pb","xhi","x","o"})
	 fmt.Println(v)
}

//"OgEu755"
//["enough","these","play","wide","wonder","box","arrive","money","tax","thus"]
func TestShortestCompletingWord (t *testing.T) {
	v := shortestCompletingWord("OgEu755", []string{"enough","these","play","wide","wonder","box","arrive","money","tax","thus"})
	fmt.Println("============")
	fmt.Println(v)
}

func TestIsAlienSorted(t *testing.T) {
/*	words := []string{"apple","app"}
	order := "worldabcefghijkmnpqstuvxyz"
*/

/*	words := []string{"fxasxpc","dfbdrifhp","nwzgs","cmwqriv","ebulyfyve","miracx","sxckdwzv","dtijzluhts","wwbmnge","qmjwymmyox"}
	order := "zkgwaverfimqxbnctdplsjyohu"
*/

/*	words := []string{"hello","leetcode"}
	order := "hlabcdefgijkmnopqrstuvwxyz"*/

	/*words := []string{"kuvp","q"}
	order := "ngxlkthsjuoqcpavbfdermiywz"*/

	words := []string{"dthkre","vvjrym","zvrzapstvm","dcjkqyrr","ditmpogp","gpzzmilb","dymhwov","bytjqyqupf","vjeaumrdy","lwvozpyx"}
	order := "tngjwhflzomueqkisxcpybdavr"

	v := isAlienSorted(words, order)
	fmt.Println(v)
}
