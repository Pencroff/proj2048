package rng

/*
#include <stdio.h>
#include <stdlib.h>
#include <inttypes.h>

#if __SIZEOF_INT128__
    typedef __uint128_t pcg128_t;
    #define PCG_128BIT_CONSTANT(high,low) \
            ((((pcg128_t)high) << 64) + low)
    #define PCG_HAS_128BIT_OPS 1
#endif

#define PCG_DEFAULT_MULTIPLIER_128 \
        PCG_128BIT_CONSTANT(2549297995355413924ULL,4865540595714422341ULL)

typedef struct pcg_state_setseq_128     pcg64_random_t;

struct pcg_state_setseq_128 {
    pcg128_t state;
    pcg128_t inc;
};

#define pcg64_srandom_r                 pcg_setseq_128_srandom_r
#define pcg64_random_r                  pcg_setseq_128_xsl_rr_64_random_r

pcg128_t mul;

inline void pcg_setseq_128_step_r(struct pcg_state_setseq_128* rng)
{
    rng->state = rng->state * PCG_DEFAULT_MULTIPLIER_128 + rng->inc;
}

inline void pcg_setseq_128_srandom_r(struct pcg_state_setseq_128* rng,
                                     pcg128_t initstate, pcg128_t initseq)
{
    rng->state = 0U;
    rng->inc = (initseq << 1u) | 1u;
    pcg_setseq_128_step_r(rng);
    rng->state += initstate;
    pcg_setseq_128_step_r(rng);
}

inline uint64_t pcg_rotr_64(uint64_t value, unsigned int rot)
{
	return (value >> rot) | (value << ((- rot) & 63));
}

inline uint64_t pcg_output_xsl_rr_128_64(pcg128_t state)
{
    return pcg_rotr_64(((uint64_t)(state >> 64u)) ^ (uint64_t)state,
                       state >> 122u);
}

inline uint64_t
pcg_setseq_128_xsl_rr_64_random_r(struct pcg_state_setseq_128* rng)
{
    pcg_setseq_128_step_r(rng);
    return pcg_output_xsl_rr_128_64(rng->state);
}

pcg64_random_t init() {
	pcg64_random_t rng;
	pcg64_srandom_r(&rng, 42u, 54u);
	mul = PCG_DEFAULT_MULTIPLIER_128;
	return rng;
}

void seed (struct pcg_state_setseq_128* rng, uint64_t value) {
	pcg128_t v = PCG_128BIT_CONSTANT(0, value);
	pcg64_srandom_r(rng, v, 54u);
}
*/
import "C"

type PcgCRng struct {
	rng               C.pcg64_random_t
	defaultMultiplier C.pcg128_t
	floatMul          float64
}

func NewPcgCRng() *PcgCRng {
	v := &PcgCRng{
		rng:               C.init(),
		defaultMultiplier: C.mul,
		floatMul:          1 / float64(1<<64-1),
	}
	v.Seed(11111)
	return v
}

func (r *PcgCRng) Seed(v int64) {
	C.seed(&r.rng, C.uint64_t(v))
}

func (r *PcgCRng) Uint64() uint64 {
	return uint64(C.pcg64_random_r(&r.rng))
}

func (r *PcgCRng) Float64() float64 {
	rnd := r.Uint64()
	return float64(rnd) * r.floatMul
}
