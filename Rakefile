require 'rake'
require 'rake/clean'
require 'cucumber'
require 'cucumber/rake/task'
require 'rspec/core/rake_task'
require 'bundler/gem_tasks'

CLEAN.include('pkg/', 'tmp/')


RSpec::Core::RakeTask.new(:spec)
Cucumber::Rake::Task.new(:cucumber)

task default: :cucumber
