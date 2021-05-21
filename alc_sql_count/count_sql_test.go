package alc_sql_count

import (
	"fmt"
	"testing"
)

const sql = `select d.no 
, d.core_data->'$.d_ah9dhsq7l2ps' as d_ah9dhsq7l2ps 
, (select title from module_data where model_no='af4218uua5fk' and no=d.core_data->'$.r_aeeme7cbx1c0') as r_aeeme7cbx1c0 
, d.core_data->'$.d_aeagalx6639c' as d_aeagalx6639c , d.core_data->'$.d_aeagm912cs8w' as d_aeagm912cs8w 
, (select title from module_data where model_no='agglmnasl81s' and no=d.core_data->'$.r_ajwtgh85do1s') as r_ajwtgh85do1s 
, d.core_data->'$.d_aeenlvme2ry8' as d_aeenlvme2ry8 , d.created_at , d.core_data->'$.d_aeempykt2ark' as d_aeempykt2ark 
, d.core_data->'$.d_aeagtx32vpc0' as d_aeagtx32vpc0 , (select username from member m where m.no=d.updated_by limit 1) as updated_by 
from module_data d where d.deleted_at is null and d.model_no = 'aeadsqwv7n5s' 
and id in (select title from module_data where model_no='agglmnasl81s' and no=d.core_data->'$.r_ajwtgh85do1s') 
group by d.id order by d.id desc`

func TestA(t *testing.T) {
	fmt.Println(Convert(sql))
}
func TestB(t *testing.T) {
	fmt.Println(Convert("select m.no ,m.title ,m.info ,m.meta_type from meta m where m.deleted_at is null and m.team_no = 'tb' order by m.id desc"))
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Convert(sql)
	}
}
