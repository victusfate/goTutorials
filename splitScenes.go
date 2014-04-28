package main

import(
    "fmt"
)

type Bounds struct {
    start int
    end int
}

type Script struct {
    length int

}

func (this *Script) splitScenes(iMaxSize int) []Bounds {
    iLength := this.length;
    iRemaining := iLength % iMaxSize
    var aScenes []Bounds
    fmt.Println("0",aScenes)
    

    if (iLength > iMaxSize) {
        for iStart := 0; iStart < iLength; iStart += iMaxSize {
            iEnd := iStart + iMaxSize - 1;

            if (iEnd >= iLength) {
                break;
            }
            aScenes = append(aScenes,Bounds{iStart,iEnd})
            fmt.Println("1",aScenes)
        }

        if (iRemaining > 0) {
            if ( float32(iRemaining) < (float32(iMaxSize) * .25)) {  // Remaining is less than 10% of scene size, just add it to the end
                aScenes[len(aScenes) - 1].end += iRemaining
                fmt.Println("2",aScenes)
            } else {
                aScenes = append(aScenes, Bounds{ iLength - iRemaining, iLength - 1 })
                fmt.Println("3",aScenes)
            }
        }
    } else {
        aScenes = append(aScenes, Bounds{ 0, iLength })
        fmt.Println("4",aScenes)
    }

    return aScenes;

}

func main() {
    s := Script{length: 1407}
    theBounds := s.splitScenes(240)
    for i,aBound := range theBounds {
        fmt.Println(i,aBound)
    }
}



// original js method
// Script.prototype.splitScenes = function(iMaxSize) {
//     var iLength    = this.getLength();
//     var iRemaining = iLength % iMaxSize;

//     var aScenes    = [];

//     if (iLength > iMaxSize) {
//         for (var iStart = 0; iStart < iLength; iStart += iMaxSize) {
//             var iEnd = iStart + iMaxSize - 1;

//             if (iEnd >= iLength) {
//                 break;
//             }

//             aScenes.push({
//                 start: iStart,
//                 end:   iEnd
//             });
//         }

//         if (iRemaining) {
//             if (iRemaining < (iMaxSize * .25)) {  // Remaining is less than 10% of scene size, just add it to the end
//                 aScenes[aScenes.length - 1].end += iRemaining;
//             } else {
//                 aScenes.push({
//                     start: iLength - iRemaining,
//                     end:   iLength - 1
//                 });
//             }
//         }
//     } else {
//         aScenes.push({
//             start: 0,
//             end:   iLength
//         });
//     }

//     return aScenes;
// };