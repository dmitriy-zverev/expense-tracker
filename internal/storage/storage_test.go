package storage

import (
	"os"
	"testing"
)

func TestGetFileData(t *testing.T) {
	setupTestData(t)
	defer cleanupTestData(t)

	tests := []struct {
		name     string
		fileName string
		setup    func() error
		want     string
		wantErr  bool
	}{
		{
			name:     "Read existing file",
			fileName: "./test_data/test.txt",
			setup: func() error {
				return os.WriteFile("./test_data/test.txt", []byte("test content"), 0755)
			},
			want:    "test content",
			wantErr: false,
		},
		{
			name:     "Read non-existent file (should create empty)",
			fileName: "./test_data/new.txt",
			setup:    func() error { return nil },
			want:     "",
			wantErr:  false,
		},
		{
			name:     "Read empty file",
			fileName: "./test_data/empty.txt",
			setup: func() error {
				return os.WriteFile("./test_data/empty.txt", []byte(""), 0755)
			},
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.setup(); err != nil {
				t.Fatalf("Setup failed: %v", err)
			}

			data, err := GetFileData(tt.fileName)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if string(data) != tt.want {
				t.Errorf("GetFileData() = %v, want %v", string(data), tt.want)
			}

			// Verify file was created if it didn't exist
			if _, err := os.Stat(tt.fileName); os.IsNotExist(err) {
				t.Errorf("File should have been created: %v", tt.fileName)
			}
		})
	}
}

func TestWriteFileData(t *testing.T) {
	setupTestData(t)
	defer cleanupTestData(t)

	tests := []struct {
		name     string
		fileName string
		data     []byte
		wantErr  bool
	}{
		{
			name:     "Write to new file",
			fileName: "./test_data/write_test.txt",
			data:     []byte("test data"),
			wantErr:  false,
		},
		{
			name:     "Write to existing file (overwrite)",
			fileName: "./test_data/existing.txt",
			data:     []byte("new content"),
			wantErr:  false,
		},
		{
			name:     "Write empty data",
			fileName: "./test_data/empty_write.txt",
			data:     []byte(""),
			wantErr:  false,
		},
		{
			name:     "Write binary data",
			fileName: "./test_data/binary.dat",
			data:     []byte{0x00, 0x01, 0x02, 0xFF},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create existing file for overwrite test
			if tt.name == "Write to existing file (overwrite)" {
				os.WriteFile(tt.fileName, []byte("old content"), 0755)
			}

			err := WriteFileData(tt.fileName, tt.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("WriteFileData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verify file was written correctly
				readData, err := os.ReadFile(tt.fileName)
				if err != nil {
					t.Errorf("Failed to read written file: %v", err)
				}

				if string(readData) != string(tt.data) {
					t.Errorf("Written data = %v, want %v", string(readData), string(tt.data))
				}

				// Check file permissions (on macOS, permissions may be different due to umask)
				info, err := os.Stat(tt.fileName)
				if err != nil {
					t.Errorf("Failed to stat file: %v", err)
				}
				// Just verify file is readable and writable by owner
				mode := info.Mode().Perm()
				if mode&0600 != 0600 {
					t.Errorf("File should be readable and writable by owner, got %v", mode)
				}
			}
		})
	}
}

func TestAppendFileData(t *testing.T) {
	setupTestData(t)
	defer cleanupTestData(t)

	tests := []struct {
		name        string
		fileName    string
		initialData string
		appendData  []byte
		want        string
		wantErr     bool
	}{
		{
			name:        "Append to existing file",
			fileName:    "./test_data/append_test.txt",
			initialData: "initial content",
			appendData:  []byte(" appended"),
			want:        "initial content appended",
			wantErr:     false,
		},
		{
			name:        "Append to new file",
			fileName:    "./test_data/new_append.txt",
			initialData: "",
			appendData:  []byte("first content"),
			want:        "first content",
			wantErr:     false,
		},
		{
			name:        "Append empty data",
			fileName:    "./test_data/empty_append.txt",
			initialData: "existing",
			appendData:  []byte(""),
			want:        "existing",
			wantErr:     false,
		},
		{
			name:        "Multiple appends",
			fileName:    "./test_data/multi_append.txt",
			initialData: "start",
			appendData:  []byte(" middle"),
			want:        "start middle",
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup initial file content if needed
			if tt.initialData != "" {
				err := os.WriteFile(tt.fileName, []byte(tt.initialData), 0755)
				if err != nil {
					t.Fatalf("Failed to setup initial file: %v", err)
				}
			}

			err := AppendFileData(tt.fileName, tt.appendData)

			if (err != nil) != tt.wantErr {
				t.Errorf("AppendFileData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verify file content
				readData, err := os.ReadFile(tt.fileName)
				if err != nil {
					t.Errorf("Failed to read file after append: %v", err)
				}

				if string(readData) != tt.want {
					t.Errorf("File content after append = %v, want %v", string(readData), tt.want)
				}

				// Test multiple appends for the multi_append test
				if tt.name == "Multiple appends" {
					err = AppendFileData(tt.fileName, []byte(" end"))
					if err != nil {
						t.Errorf("Second append failed: %v", err)
					}

					readData, err = os.ReadFile(tt.fileName)
					if err != nil {
						t.Errorf("Failed to read file after second append: %v", err)
					}

					expected := "start middle end"
					if string(readData) != expected {
						t.Errorf("File content after second append = %v, want %v", string(readData), expected)
					}
				}
			}
		})
	}
}

func TestCreateIfNotCreated(t *testing.T) {
	setupTestData(t)
	defer cleanupTestData(t)

	tests := []struct {
		name     string
		fileName string
		setup    func() error
		wantErr  bool
	}{
		{
			name:     "Create new file",
			fileName: "./test_data/new_file.txt",
			setup:    func() error { return nil },
			wantErr:  false,
		},
		{
			name:     "File already exists",
			fileName: "./test_data/existing_file.txt",
			setup: func() error {
				return os.WriteFile("./test_data/existing_file.txt", []byte("existing"), 0755)
			},
			wantErr: false,
		},
		{
			name:     "Create in nested directory",
			fileName: "./test_data/nested/dir/file.txt",
			setup: func() error {
				return os.MkdirAll("./test_data/nested/dir", 0755)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.setup(); err != nil {
				t.Fatalf("Setup failed: %v", err)
			}

			err := createIfNotCreated(tt.fileName)

			if (err != nil) != tt.wantErr {
				t.Errorf("createIfNotCreated() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verify file exists
				if _, err := os.Stat(tt.fileName); os.IsNotExist(err) {
					t.Errorf("File should exist after createIfNotCreated: %v", tt.fileName)
				}

				// For existing file test, verify content wasn't changed
				if tt.name == "File already exists" {
					data, err := os.ReadFile(tt.fileName)
					if err != nil {
						t.Errorf("Failed to read existing file: %v", err)
					}
					if string(data) != "existing" {
						t.Errorf("Existing file content changed: got %v, want 'existing'", string(data))
					}
				}
			}
		})
	}
}

func TestFileOperationsIntegration(t *testing.T) {
	setupTestData(t)
	defer cleanupTestData(t)

	fileName := "./test_data/integration_test.txt"

	// Test complete workflow: write, read, append, read
	t.Run("Complete file operations workflow", func(t *testing.T) {
		// 1. Write initial data
		initialData := []byte("initial content")
		err := WriteFileData(fileName, initialData)
		if err != nil {
			t.Fatalf("WriteFileData failed: %v", err)
		}

		// 2. Read the data
		readData, err := GetFileData(fileName)
		if err != nil {
			t.Fatalf("GetFileData failed: %v", err)
		}
		if string(readData) != string(initialData) {
			t.Errorf("Read data = %v, want %v", string(readData), string(initialData))
		}

		// 3. Append more data
		appendData := []byte(" + appended content")
		err = AppendFileData(fileName, appendData)
		if err != nil {
			t.Fatalf("AppendFileData failed: %v", err)
		}

		// 4. Read final data
		finalData, err := GetFileData(fileName)
		if err != nil {
			t.Fatalf("Final GetFileData failed: %v", err)
		}
		expected := "initial content + appended content"
		if string(finalData) != expected {
			t.Errorf("Final data = %v, want %v", string(finalData), expected)
		}

		// 5. Overwrite with new data
		newData := []byte("completely new content")
		err = WriteFileData(fileName, newData)
		if err != nil {
			t.Fatalf("Overwrite WriteFileData failed: %v", err)
		}

		// 6. Verify overwrite
		overwriteData, err := GetFileData(fileName)
		if err != nil {
			t.Fatalf("Overwrite GetFileData failed: %v", err)
		}
		if string(overwriteData) != string(newData) {
			t.Errorf("Overwrite data = %v, want %v", string(overwriteData), string(newData))
		}
	})
}

func TestErrorCases(t *testing.T) {
	setupTestData(t)
	defer cleanupTestData(t)

	t.Run("WriteFileData to invalid path", func(t *testing.T) {
		// Try to write to a path that doesn't exist and can't be created
		invalidPath := "/invalid/path/that/cannot/be/created/file.txt"
		err := WriteFileData(invalidPath, []byte("test"))
		if err == nil {
			t.Errorf("WriteFileData should fail for invalid path")
		}
	})

	t.Run("AppendFileData to invalid path", func(t *testing.T) {
		// Try to append to a path that doesn't exist and can't be created
		invalidPath := "/invalid/path/that/cannot/be/created/file.txt"
		err := AppendFileData(invalidPath, []byte("test"))
		if err == nil {
			t.Errorf("AppendFileData should fail for invalid path")
		}
	})

	t.Run("GetFileData from invalid path", func(t *testing.T) {
		// Try to read from a path that doesn't exist and can't be created
		invalidPath := "/invalid/path/that/cannot/be/created/file.txt"
		_, err := GetFileData(invalidPath)
		if err == nil {
			t.Errorf("GetFileData should fail for invalid path")
		}
	})
}

func TestAppendFileDataErrorHandling(t *testing.T) {
	setupTestData(t)
	defer cleanupTestData(t)

	t.Run("AppendFileData file close on error", func(t *testing.T) {
		// Create a file first
		fileName := "./test_data/append_error_test.txt"
		err := os.WriteFile(fileName, []byte("initial"), 0755)
		if err != nil {
			t.Fatalf("Failed to create initial file: %v", err)
		}

		// Make the file read-only to cause write error
		err = os.Chmod(fileName, 0444)
		if err != nil {
			t.Fatalf("Failed to change file permissions: %v", err)
		}

		// Try to append - this should fail but not panic
		err = AppendFileData(fileName, []byte("append"))
		if err == nil {
			t.Errorf("AppendFileData should fail for read-only file")
		}

		// Restore permissions for cleanup
		os.Chmod(fileName, 0755)
	})
}

func TestEdgeCases(t *testing.T) {
	setupTestData(t)
	defer cleanupTestData(t)

	t.Run("Large file operations", func(t *testing.T) {
		fileName := "./test_data/large_file.txt"

		// Create large data (1MB)
		largeData := make([]byte, 1024*1024)
		for i := range largeData {
			largeData[i] = byte(i % 256)
		}

		// Write large file
		err := WriteFileData(fileName, largeData)
		if err != nil {
			t.Errorf("WriteFileData failed for large file: %v", err)
		}

		// Read large file
		readData, err := GetFileData(fileName)
		if err != nil {
			t.Errorf("GetFileData failed for large file: %v", err)
		}

		if len(readData) != len(largeData) {
			t.Errorf("Large file size mismatch: got %d, want %d", len(readData), len(largeData))
		}
	})

	t.Run("File with special characters in name", func(t *testing.T) {
		fileName := "./test_data/file with spaces & special-chars.txt"
		testData := []byte("special file content")

		err := WriteFileData(fileName, testData)
		if err != nil {
			t.Errorf("WriteFileData failed for special filename: %v", err)
		}

		readData, err := GetFileData(fileName)
		if err != nil {
			t.Errorf("GetFileData failed for special filename: %v", err)
		}

		if string(readData) != string(testData) {
			t.Errorf("Special filename content mismatch: got %v, want %v", string(readData), string(testData))
		}
	})
}

// Helper functions for test setup
func setupTestData(t *testing.T) {
	// Create test data directory
	err := os.MkdirAll("./test_data", 0755)
	if err != nil {
		t.Fatalf("Failed to create test data directory: %v", err)
	}
}

func cleanupTestData(t *testing.T) {
	err := os.RemoveAll("./test_data")
	if err != nil {
		t.Logf("Failed to cleanup test data: %v", err)
	}
}
