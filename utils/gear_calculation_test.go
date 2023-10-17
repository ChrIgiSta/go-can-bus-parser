/**
 * Copyright © 2023, Staufi Tech - Switzerland
 * All rights reserved.
 *
 *   ________________________   ___ _     ________________  _  ____
 *  / _____  _  ____________/  / __|_|   /_______________  | | ___/
 * ( (____ _| |_ _____ _   _ _| |__ _      | |_____  ____| |_|_
 *  \____ (_   _|____ | | | (_   __) |     | | ___ |/ ___)  _  \
 *  _____) )| |_/ ___ | |_| | | |  | |     | | ____( (___| | | |
 * (______/  \__)_____|____/  |_|  |_|     |_|_____)\____)_| |_|
 *
 *
 *  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 *  AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 *  IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 *  ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
 *  LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 *  CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 *  SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 *  INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 *  CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 *  ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 *  POSSIBILITY OF SUCH DAMAGE.
 */

package utils

import "testing"

func TestFirstGear(t *testing.T) {
	gc := NewGearCalculator(AstraHOpcGears, GEAR_CALC_RATIO_TO_RPM_PER_KMH, 9)
	g, err := gc.Get(1416, 12)
	if err != nil {
		t.Error(err)
	}
	if g.Gear != 1 {
		t.Error("should be the first gear")
	}

	g, err = gc.Get(3220, 28)
	if err != nil {
		t.Error(err)
	}
	if g.Gear != 1 {
		t.Error("should be the first gear")
	}
}

func TestSecondGear(t *testing.T) {
	gc := NewGearCalculator(AstraHOpcGears, GEAR_CALC_RATIO_TO_RPM_PER_KMH, 9)
	g, err := gc.Get(2000, 31)
	if err != nil {
		t.Error(err)
	}
	if g.Gear != 2 {
		t.Error("should be the second gear")
	}

	g, err = gc.Get(3000, 48)
	if err != nil {
		t.Error(err)
	}
	if g.Gear != 2 {
		t.Error("should be the second gear")
	}
}

func TestThirdGear(t *testing.T) {
	gc := NewGearCalculator(AstraHOpcGears, GEAR_CALC_RATIO_TO_RPM_PER_KMH, 8.5)
	g, err := gc.Get(2100, 50)
	if err != nil {
		t.Error(err)
	}
	if g.Gear != 3 {
		t.Error("should be the third gear")
	}

	g, err = gc.Get(3000, 70)
	if err != nil {
		t.Error(err)
	}
	if g.Gear != 3 {
		t.Error("should be the third gear")
	}
}

func TestFourthGear(t *testing.T) {
	gc := NewGearCalculator(AstraHOpcGears, GEAR_CALC_RATIO_TO_RPM_PER_KMH, 8.5)
	g, err := gc.Get(2500, 80)
	if err != nil {
		t.Error(err)
	}
	if g.Gear != 4 {
		t.Error("should be the fourth gear")
	}

	g, err = gc.Get(2900, 88)
	if err != nil {
		t.Error(err)
	}
	if g.Gear != 4 {
		t.Error("should be the fourth gear")
	}
}

func TestFifthGear(t *testing.T) {
	gc := NewGearCalculator(AstraHOpcGears, GEAR_CALC_RATIO_TO_RPM_PER_KMH, 8.5)
	g, err := gc.Get(2500, 90)
	if err != nil {
		t.Error(err)
	}
	if g.Gear != 5 {
		t.Error("should be the fifth gear")
	}

	g, err = gc.Get(2900, 105)
	if err != nil {
		t.Error(err)
	}
	if g.Gear != 5 {
		t.Error("should be the fifth gear")
	}
}

func TestSixthGear(t *testing.T) {
	gc := NewGearCalculator(AstraHOpcGears, GEAR_CALC_RATIO_TO_RPM_PER_KMH, 8.5)
	g, err := gc.Get(2250, 100)
	if err != nil {
		t.Error(err)
	}
	if g.Gear != 6 {
		t.Error("should be the sixth gear")
	}

	g, err = gc.Get(2750, 120)
	if err != nil {
		t.Error(err)
	}
	if g.Gear != 6 {
		t.Error("should be the sixth gear")
	}
}

// func TestReverseGear(t *testing.T) {
// 	gc := NewGearCalculator(AstraHOpcGears, GEAR_CALC_RATIO_TO_RPM_PER_KMH, 0)
// 	g, err := gc.Get(1100.5, 10)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if g.Gear != 6 {
// 		t.Error("should be the reverse gear")
// 	}
// }

func TestNeutralGear(t *testing.T) {
	gc := NewGearCalculator(AstraHOpcGears, GEAR_CALC_RATIO_TO_RPM_PER_KMH, 8.5)
	g, err := gc.Get(2250, 0)
	if err != nil {
		t.Error(err)
	}
	if g.Gear != 0 {
		t.Error("should be the neutral gear (no gear)")
	}
}
