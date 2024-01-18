package timeSpliceFile

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const SplitIntervalMinute string = "2006-01-02-15-04"
const SplitIntervalHour string = "2006-01-02-15"
const SplitIntervalDay string = time.DateOnly
const SplitIntervalMonth string = "2006-01"
const SplitIntervalYear string = "2006"

type SplitFile struct {
	splitInterval string
	fileDir       string
	typeSuffix    string
	filename      string
	file          *os.File
	mu            sync.Mutex
}

func New(dir, splitFileTime, typeSuffix string) (*SplitFile, error) {
	timeName := time.Now().Format(splitFileTime)
	filename := fmt.Sprintf("%s.%s", timeName, typeSuffix)
	if err := os.MkdirAll(dir, 0777); err != nil {
		return nil, err
	}
	file, err := os.OpenFile(filepath.Join(dir, filename), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &SplitFile{splitInterval: splitFileTime, fileDir: dir, filename: filename, file: file, typeSuffix: typeSuffix}, nil
}

func (l *SplitFile) Write(p []byte) (n int, err error) {
	timeName := time.Now().Format(l.splitInterval)
	filename := fmt.Sprintf("%s.%s", timeName, l.typeSuffix)
	if l.filename != filename {
		l.mu.Lock()
		err := l.file.Close()
		if err != nil {
			return 0, err
		}
		file, err := os.OpenFile(filepath.Join(l.fileDir, filename), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return 0, err
		}
		l.file = file
		l.filename = filename
		l.mu.Unlock()
	}
	return l.file.Write(p)
}

func (l *SplitFile) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.file.Close()
}
