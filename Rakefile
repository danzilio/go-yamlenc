require 'rake'
require 'rake/clean'
require 'cucumber'
require 'cucumber/rake/task'

CLEAN.include('pkg/', 'tmp/')

Cucumber::Rake::Task.new(:cucumber)

source_files = Dir['*.go']

task :build do
  sh "go build -o bin/enc #{source_files.join ' '}"
end

task :test do
  sh "go test"
  Rake::Task[:build].execute
  Rake::Task[:cucumber].execute
end

task default: :test
