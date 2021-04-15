/* Programmers coding test
* 트럭 여러 대가 강을 가로지르는 일 차선 다리를 정해진 순으로 건너려 합니다.
* 모든 트럭이 다리를 건너려면 최소 몇 초가 걸리는지 알아내시오.
* 트럭은 1초 1만큼 움직이며, 다리 길이는 bridge_length, 다리 무게는 weight까지 견딥니다.
*
* solution 함수의 매개 변수로 다리 길이 bridge_length,
* 다리가 견딜 수 있는 무게 weight, 트러별 무게 truck_weights가 주어집니다.
* 이때 모든 트럭이 다리를 건너려면 최소 몇 초가 걸리는지 return 하는 solution 함수를 완성합니다.
 */

package main

import "fmt"

type truck struct {
	weight int
	curLen int // 트럭의 이동 거리
}

// 다리 구조체
type bridge struct {
	length    int
	weight    int
	curWeight int // 현재 다리를 지나고 있는 트럭의 총 무게
	trucks    []truck
}

func solution(bridge_length int, weight int, truck_weights []int) int {
	// 다리 생성
	br := bridge{length: bridge_length, weight: weight, curWeight: 0}

	// 처음 대기를 제외한 시작부터 time을 측정하기 위해 1로 시작
	time := 1
	// 다리에 트럭이 없고 대기 중인 트럭도 없을 경우 종료
	for (len(truck_weights) > 0) || (len(br.trucks) > 0) {
		// 다리에 트럭이 있을 경우 동작
		if len(br.trucks) > 0 {
			time += 1
			for i := 0; i < len(br.trucks); i++ {
				br.trucks[i].curLen -= 1
			}
			if br.trucks[0].curLen == 0 { // 트럭의 이동거리가 0일 경우 도착으로 간주
				br.curWeight -= br.trucks[0].weight // queue 구조로 first truck을 pop
				br.trucks = br.trucks[1:]
			}
		}
		if len(truck_weights) > 0 {
			// 다리가 견딜 수 있는 무게를 초과 시 트럭은 대기
			if br.weight >= truck_weights[0]+br.curWeight {
				var truckWeight int
				truckWeight, truck_weights = truck_weights[0], truck_weights[1:]
				br.curWeight += truckWeight
				br.trucks = append(br.trucks, truck{truckWeight, br.length})
			}
		}

		fmt.Println(time, "|", br.curWeight, "|", br.trucks, "|", truck_weights)
	}

	return time
}

func main() {
	var truck_weights []int
	truck_weights = append(truck_weights, 7)
	truck_weights = append(truck_weights, 4)
	truck_weights = append(truck_weights, 5)
	truck_weights = append(truck_weights, 6)

	fmt.Println(solution(2, 10, truck_weights))
}
