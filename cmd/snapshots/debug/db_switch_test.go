package debug

import (
	"context"
	"github.com/anacrolix/torrent/bencode"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ledgerwatch/turbo-geth/common/dbutils"
	"github.com/ledgerwatch/turbo-geth/core/rawdb"
	"github.com/ledgerwatch/turbo-geth/ethdb"
	"github.com/ledgerwatch/turbo-geth/log"
	trnt "github.com/ledgerwatch/turbo-geth/turbo/snapshotsync/bittorrent"
	"os"
	"testing"
)

func TestDbSwitch(t *testing.T) {
	snapshotPath:="/media/b00ris/nvme/tmp/canonicalswitch"
	dbPath:="/media/b00ris/nvme/fresh_sync/tg/chaindata/"
	toBlock:=uint64(5000)
	epoch:=1000
	err := os.RemoveAll(snapshotPath)
	if err != nil {
		t.Fatal(err)
	}
	kv := ethdb.NewLMDB().Path(dbPath).MustOpen()

	snKV := ethdb.NewLMDB().WithBucketsConfig(func(defaultBuckets dbutils.BucketsCfg) dbutils.BucketsCfg {
		return dbutils.BucketsCfg{
			dbutils.HeaderPrefix:              dbutils.BucketConfigItem{},
		}
	}).Path(snapshotPath).MustOpen()

	snKV:=ethdb.NewSnapshot2KV().DB(snKV).MustOpen()

	db := ethdb.NewObjectDatabase(kv)
	snDB := ethdb.NewObjectDatabase(snKV)
	tx,err:=snDB.Begin(context.Background(), ethdb.RW)
	if err!=nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	var hash common.Hash
	var header []byte
	for i := uint64(1); i <= toBlock; i++ {
		hash, err = rawdb.ReadCanonicalHash(db, i)
		if err != nil {
			t.Fatal(err)
		}
		header = rawdb.ReadHeaderRLP(db, hash, i)
		if len(header) == 0 {
			t.Fatal(err)
		}

		err = tx.Append(dbutils.HeaderPrefix, dbutils.HeaderKey(i, hash), header)
		if err != nil {
			t.Fatal(err)
		}
		if i%uint64(epoch) == 0 {
			_, err=tx.Commit()
			if err!=nil {
				t.Fatal(err)
			}
			snKVSwitch := ethdb.NewLMDB().WithBucketsConfig(func(defaultBuckets dbutils.BucketsCfg) dbutils.BucketsCfg {
				return dbutils.BucketsCfg{
					dbutils.HeaderPrefix:              dbutils.BucketConfigItem{},
				}
			}).Path(snapshotPath).MustOpen()
			snDB.KV().(*ethdb.SnapshotKV2).DbSwitch(snKVSwitch)

		}
	}
	tx.Rollback()
	snDB.Close()
	err = os.Remove(snapshotPath + "/lock.mdb")
	if err != nil {
		log.Warn("Remove lock", "err", err)
		t.Fatal(err)
	}
	err = os.Remove(snapshotPath + "/LOCK")
	if err != nil {
		log.Warn("Remove lock", "err", err)
		t.Fatal(err)
	}

	info, err := trnt.BuildInfoBytesForLMDBSnapshot(snapshotPath,trnt.LmdbFilename)
	if err != nil {
		t.Fatal(err)
	}
	infoBytes1, err := bencode.Marshal(info)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(metainfo.HashBytes(infoBytes1))
}